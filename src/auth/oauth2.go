package auth

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	return &oauth2Service{
		Context: context,
		Config:  config,
	}
}

func NewOAuth2ServiceFactory(
	config *oauth2.Config,
) OAuth2ServiceFactory {
	return &oauth2ServiceFactory{
		Config: config,
	}
}

type oauth2Service struct {
	Context *gin.Context
	Config  *oauth2.Config
	token   *oauth2.Token
	err     error
}

func (service *oauth2Service) Token() (*oauth2.Token, error) {
	service.compareState()
	service.exchangeCode()
	return service.token, service.err
}

func (service *oauth2Service) compareState() {
	queryState := service.queryState()
	sessionState := service.sessionState()
	service.compare(queryState, sessionState)
}

func (service *oauth2Service) compare(query, session string) {
	if service.hasError() {
		return
	}
	if query != session {
		service.err = ErrStateNotMatch
	}
}

func (service *oauth2Service) queryState() string {
	return service.query(stateKey)
}

func (service *oauth2Service) sessionState() string {
	return service.sessionValue(stateKey)
}

func (service *oauth2Service) exchangeCode() {
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

func (service *oauth2Service) queryCode() string {
	return service.query(codeKey)
}

func (service *oauth2Service) sessionVerifier() string {
	return service.sessionValue(verifierKey)
}

func (service *oauth2Service) exchangeContext() context.Context {
	return context.TODO()
}

func (service *oauth2Service) exchangeOptions() (opts []oauth2.AuthCodeOption) {
	opts = make([]oauth2.AuthCodeOption, 0)

	verifier := service.sessionVerifier()
	opts = append(opts, oauth2.SetAuthURLParam(codeVerifierKey, verifier))
	return
}

func (service *oauth2Service) query(key string) string {
	return service.Context.Query(key)
}

func (service *oauth2Service) sessionValue(key string) string {
	s := service.session()
	value := s.Get(key)
	if value == nil {
		service.err = errors.Wrap(ErrSessionValueNotFound, key+"value is not found")
		return ""
	}
	return value.(string)
}

func (service *oauth2Service) hasError() bool {
	return service.err != nil
}

func (service *oauth2Service) session() sessions.Session {
	return sessions.Default(service.Context)
}

type oauth2ServiceFactory struct {
	Config *oauth2.Config
}

func (factory *oauth2ServiceFactory) NewOAuth2Service(context *gin.Context) OAuth2Service {
	return NewOAuth2Service(
		context,
		factory.Config,
	)
}
