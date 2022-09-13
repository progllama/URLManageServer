package auth

import (
	"context"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OAuth2Service interface {
	Token() (*oauth2.Token, error)
}

type OAuth2ServiceFactory interface {
	NewOAuth2Service(*gin.Context) OAuth2Service
}

func NewOAuth2Service(
	context *gin.Context,
	config *oauth2.Config,
) OAuth2Service {
	return &oAuth2Service{
		Context: context,
		Config:  config,
	}
}

type oAuth2Service struct {
	Context *gin.Context
	Config  *oauth2.Config
	token   *oauth2.Token
	err     error
}

func (service *oAuth2Service) Token() (*oauth2.Token, error) {
	service.compareState()
	service.exchangeCode()
	return service.token, service.err
}

func (service *oAuth2Service) compareState() {
	queryState := service.queryState()
	sessionState := service.sessionState()
	service.compare(queryState, sessionState)
}

func (service *oAuth2Service) compare(query, session string) {
	if service.hasError() {
		return
	}
	if query != session {
		service.err = errors.New("state does not match")
	}
}

func (service *oAuth2Service) queryState() string {
	return service.query(stateKey)
}

func (service *oAuth2Service) sessionState() string {
	return service.sessionValue(stateKey)
}

func (service *oAuth2Service) exchangeCode() {
	if service.hasError() {
		return
	}

	conf := service.Config
	context := service.exchangeContext()
	code := service.queryCode()
	opts := service.exchangeOptions()

	service.token, service.err = conf.Exchange(
		context,
		code,
		opts...,
	)
}

func (service *oAuth2Service) queryCode() string {
	return service.query(codeKey)
}

func (service *oAuth2Service) sessionVerifier() string {
	return service.sessionValue(verifierKey)
}

func (service *oAuth2Service) exchangeContext() context.Context {
	return context.TODO()
}

func (service *oAuth2Service) exchangeOptions() (opts []oauth2.AuthCodeOption) {
	opts = make([]oauth2.AuthCodeOption, 0)

	verifier := service.sessionVerifier()
	opts = append(opts, oauth2.SetAuthURLParam(codeVerifierKey, verifier))
	return
}

func (service *oAuth2Service) query(key string) string {
	return service.Context.Query(key)
}

func (service *oAuth2Service) sessionValue(key string) string {
	s := service.session()
	value := s.Get(key)
	if value == nil {
		service.err = errors.New(key + "is empty")
		return ""
	}
	return value.(string)
}

func (service *oAuth2Service) hasError() bool {
	return service.err != nil
}

func (service *oAuth2Service) session() sessions.Session {
	return sessions.Default(service.Context)
}
