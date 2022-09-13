package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var twitterScope = []string{
	"users.read",
	"tweet.read",
}

const (
	twitterRedirectURL = "http://127.0.0.1:8080/login"
	twitterProfileURL  = "https://api.twitter.com/2/users/me"
)

var TwitterConf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  twitterRedirectURL,
	Scopes:       twitterScope,
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://twitter.com/i/oauth2/authorize",
		TokenURL: "https://api.twitter.com/2/oauth2/token",
	},
}

func TwitterOAuth2RedirectHandler(ctx *gin.Context) {
	handler := CallbackHandler{
		Context:              ctx,
		Config:               TwitterConf,
		OAuth2ServiceFactory: NewOAuth2ServiceFactory(TwitterConf),
		UserServiceFactory:   NewTwitterUserServiceFactory(),
		LoginServiceFactory:  NewLoginServiceFactory(),
	}
	handler.Handle(ctx)

	s := sessions.Default(ctx)
	log.Println(s.Get(loginIdKey))
}

type TwitterUser struct {
	ID string `json:"id"`
}

func (u *TwitterUser) LoginId() string {
	return u.ID
}

type TwitterUserServiceFactory struct {
}

func (factory *TwitterUserServiceFactory) NewUserService() UserService {
	return NewTwitterUserService()
}

func NewTwitterUserServiceFactory() UserServiceFactory {
	return &TwitterUserServiceFactory{}
}

type TwitterUserService struct {
}

func NewTwitterUserService() UserService {
	return &TwitterUserService{}
}

func (service *TwitterUserService) Fetch(token *oauth2.Token) (User, error) {
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
	var u TwitterUser
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func TwitterLoginURL(state string, codeChallenge string) string {
	return TwitterConf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
}
