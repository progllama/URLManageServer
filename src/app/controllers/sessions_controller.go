package controllers

import (
	"net/http"
	"strconv"
	"url_manager/app/forms"
	"url_manager/app/repositories"
	"url_manager/app/session"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	userRepository repositories.UserRepository
	sessionFactory session.SessionFactory
}

func NewSessionController(userRepository repositories.UserRepository) *SessionController {
	return &SessionController{userRepository, nil}
}

func (ctrl *SessionController) NewSession(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}

func (ctrl *SessionController) CreateSession(c *gin.Context) {
	var form forms.LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := ctrl.userRepository.FindByLoginId(form.LoginId)
	if err != nil {
		ctrl.redirectToSignUpPage(c)
		return
	}

	if !user.Authenticate(form.LoginId, form.Password) {
		ctrl.redirectToSignUpPage(c)
		return
	}

	ctrl.login(c, user.ID)

	c.Redirect(http.StatusFound, "/users/"+strconv.Itoa(int(user.ID)))
}

func (ctrl *SessionController) DestroySession(c *gin.Context) {
	ctrl.logout(c)
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (ctrl *SessionController) login(ctx *gin.Context, userId int) {
	s := ctrl.getNewSession(ctx)
	s.SetUserId(userId)
}

func (ctrl *SessionController) logout(ctx *gin.Context) {
	session := ctrl.getNewSession(ctx)
	session.Clear()
}

func (ctrl *SessionController) getNewSession(c *gin.Context) session.Session {
	return ctrl.sessionFactory.Create(c)
}

func (ctrl *SessionController) redirectToSignUpPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/users/new")
}
