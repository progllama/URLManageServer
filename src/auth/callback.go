package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type CallbackHandler struct {
	Context              *gin.Context
	Config               *oauth2.Config
	token                *oauth2.Token
	OAuth2ServiceFactory OAuth2ServiceFactory
	UserServiceFactory   UserServiceFactory
	LoginServiceFactory  LoginServiceFactory
	user                 User
	err                  error
}

func (handler *CallbackHandler) Handle(ctx *gin.Context) {
	handler.fetchToken()
	handler.fetchUser()
	handler.associate()
	handler.login()
	handler.setResponse()
}

func (handler *CallbackHandler) fetchToken() {
	service := handler.getNewOAuth2Service()
	handler.token, handler.err = service.Token()
}

func (handler *CallbackHandler) getNewOAuth2Service() OAuth2Service {
	return handler.OAuth2ServiceFactory.NewOAuth2Service(
		handler.Context,
	)
}

func (handler *CallbackHandler) fetchUser() {
	if handler.hasError() {
		return
	}
	service := handler.getNewUserService()
	handler.user, handler.err = service.Fetch(handler.token)
}

func (handler *CallbackHandler) getNewUserService() UserService {
	return handler.UserServiceFactory.NewUserService()
}

func (handler *CallbackHandler) associate() {
	if handler.hasError() {
		return
	}
	handler.err = AssociateAccount(handler.user, handler.token)
}

func (handler *CallbackHandler) login() {
	if handler.hasError() {
		return
	}
	service := handler.getNewLoginService()
	handler.err = service.Login(handler.user.LoginId())
}

func (handler *CallbackHandler) getNewLoginService() LoginService {
	return handler.LoginServiceFactory.NewLoginService(handler.Context)
}

func (handler *CallbackHandler) setResponse() {
	if handler.hasError() {
		handler.Context.JSON(http.StatusInternalServerError, handler.err.Error())
	} else {
		handler.Context.Status(http.StatusOK)
	}
}

func (handler *CallbackHandler) hasError() bool {
	return handler.err != nil
}
