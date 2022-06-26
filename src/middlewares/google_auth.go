// Package google provides you access to Google's OAuth2
// infrastructure. The implementation is based on this blog post:
// http://skarlso.github.io/2016/06/12/google-signin-with-go/
package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"url_manager/models"
	"url_manager/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	goauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"gorm.io/gorm"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials stores google client-ids.
type Credentials struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"secret"`
}

const (
	stateKey  = "state"
	sessionID = "ginoauth_google_session"
)

var (
	conf       *oauth2.Config
	repository repositories.UserRepository
)

func init() {
	gob.Register(goauth.Userinfo{})
}

func randToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		glog.Fatalf("[Gin-OAuth] Failed to read rand: %v", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

// Setup the authorization path
func Setup(redirectURL, credFile string, scopes []string, secret []byte) {
	var c Credentials
	file, err := os.ReadFile(credFile)
	if err != nil {
		glog.Fatalf("[Gin-OAuth] File error: %v", err)
	}
	if err := json.Unmarshal(file, &c); err != nil {
		glog.Fatalf("[Gin-OAuth] Failed to unmarshal client credentials: %v", err)
	}

	conf = &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}
}

func SetUserRepository(repo repositories.UserRepository) {
	repository = repo
}

func Session(name string, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(name, store)
}

func LoginHandler(ctx *gin.Context) {
	stateValue := randToken()
	session := sessions.Default(ctx)
	session.Set(stateKey, stateValue)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{"authorization_endpoint": GetLoginURL(stateValue)})
}

func GetLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Handle the exchange code to initiate a transport.
		session := sessions.Default(ctx)

		existingSession := session.Get(sessionID)
		if userInfo, ok := existingSession.(goauth.Userinfo); ok {
			ctx.Set("user", userInfo)
			ctx.Next()
			return
		}

		retrievedState := session.Get(stateKey)
		if retrievedState != ctx.Query(stateKey) {
			ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", retrievedState))
			return
		}

		tok, err := conf.Exchange(context.TODO(), ctx.Query("code"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to exchange code for oauth token: %w", err))
			return
		}

		oAuth2Service, err := goauth.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, tok)))
		if err != nil {
			glog.Errorf("[Gin-OAuth] Failed to create oauth service: %v", err)
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create oauth service: %w", err))
			return
		}

		userInfo, err := oAuth2Service.Userinfo.Get().Do()
		if err != nil {
			glog.Errorf("[Gin-OAuth] Failed to get userinfo for user: %v", err)
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get userinfo for user: %w", err))
			return
		}

		_, err = repository.FindByOpenID(userInfo.Id)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			repository.Add(models.User{OpenID: userInfo.Id})
		}

		if err != nil {
			glog.Errorf("[Gin-OAuth] Failed to retrieve user data: %v", err)
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to retrieve user data: %v", err))
			return
		}

		ctx.Set("user", userInfo)

		session.Set(sessionID, userInfo)
		if err := session.Save(); err != nil {
			glog.Errorf("[Gin-OAuth] Failed to save session: %v", err)
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to save session: %v", err))
			return
		}
	}
}
