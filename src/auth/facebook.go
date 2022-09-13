package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var facebookScope = []string{
	"public_profile",
}

const (
	facebookRedirectURL = "http://localhost:8080/login"
	facebookProfileURL  = "https://graph.facebook.com/v14.0/me"
)

var FacebookConf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  facebookRedirectURL,
	Scopes:       facebookScope,
	Endpoint:     facebook.Endpoint,
}

func FacebookLoginURL(state string) string {
	return FacebookConf.AuthCodeURL(state)
}

func FaceBookOAuth2RedirectHandler(ctx *gin.Context) {
	handler := CallbackHandler{
		Context:              ctx,
		Config:               FacebookConf,
		OAuth2ServiceFactory: NewOAuth2ServiceFactory(FacebookConf),
		UserServiceFactory:   NewGoogleUserServiceFactory(),
		LoginServiceFactory:  NewLoginServiceFactory(),
	}
	handler.Handle(ctx)

	s := sessions.Default(ctx)
	log.Println(s.Get(loginIdKey))
}

type FacebookUser struct {
	ID string `json:"id"`
}

func (u *FacebookUser) LoginId() string {
	return u.ID
}

type FacebookUserServiceFactory struct {
}

func (factory *FacebookUserServiceFactory) NewUserService() UserService {
	return NewFacebookUserService()
}

func NewFacebookUserServiceFactory() UserServiceFactory {
	return &FacebookUserServiceFactory{}
}

type FacebookUserService struct {
}

func NewFacebookUserService() UserService {
	return &FacebookUserService{}
}

func (service *FacebookUserService) Fetch(token *oauth2.Token) (User, error) {
	client := oauth2.NewClient(
		context.TODO(),
		oauth2.StaticTokenSource(token),
	)
	res, err := client.Get(twitterProfileURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var u FacebookUser
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
