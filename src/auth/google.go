package auth

import (
	"context"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var googleScope = []string{
	"openid",
}

const (
	googleRedirectURL = "http://localhost:8080/login"
)

var GoogleConf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  googleRedirectURL,
	Scopes:       googleScope,
	Endpoint:     google.Endpoint,
}

func GoogleLoginURL(state string, codeChallenge string) string {
	return GoogleConf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
}

func GoogleOAuth2RedirectHandler(ctx *gin.Context) {
	handler := CallbackHandler{
		Context:              ctx,
		Config:               GoogleConf,
		OAuth2ServiceFactory: NewOAuth2ServiceFactory(GoogleConf),
		UserServiceFactory:   NewGoogleUserServiceFactory(),
		LoginServiceFactory:  NewLoginServiceFactory(),
	}
	handler.Handle(ctx)

	s := sessions.Default(ctx)
	log.Println(s.Get(loginIdKey))
}

type GoogleUserServiceFactory struct {
}

func (factory *GoogleUserServiceFactory) NewUserService() UserService {
	return NewGoogleUserService()
}

func NewGoogleUserServiceFactory() UserServiceFactory {
	return &GoogleUserServiceFactory{}
}

type GoogleUserService struct {
}

func NewGoogleUserService() UserService {
	return &GoogleUserService{}
}

func (service *GoogleUserService) Fetch(token *oauth2.Token) (User, error) {
	gservice, err := goauth.NewService(
		context.TODO(),
		option.WithTokenSource(GoogleConf.TokenSource(context.TODO(), token)))
	if err != nil {
		return nil, err
	}

	userInfo, err := gservice.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}
	return &GoogleUser{ID: userInfo.Id}, nil
}

type GoogleUser struct {
	ID string `json:"id"`
}

func (user *GoogleUser) LoginId() string {
	return user.ID
}
