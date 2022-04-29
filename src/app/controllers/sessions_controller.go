package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"url_manager/app"
	"url_manager/app/forms"
	"url_manager/app/repositories"

	"github.com/gin-gonic/gin"
)

type SessionController struct{}

func NewSessionController() SessionController {
	return SessionController{}
}

func (ctrl *SessionController) NewSession(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}

func (ctrl *SessionController) CreateSession(c *gin.Context) {
	fmt.Println("try bind")
	var form forms.LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Println("success login.")

	repo := repositories.NewUserRepository()
	user, err := repo.FindByLoginId(form.LoginId)
	if err != nil {
		ctrl.redirectToSignUp(c)
		return
	}

	if !user.Authenticate(form.LoginId, form.Password) {
		ctrl.redirectToSignUp(c)
		return
	}

	session := ctrl.getNewSession(c)
	session.SetUserId(int(user.ID))
	fmt.Println("success", user.ID)

	c.Redirect(http.StatusFound, "/users/"+strconv.Itoa(int(user.ID)))
}

func (ctrl *SessionController) DestroySession(c *gin.Context) {
	session := ctrl.getNewSession(c)
	session.Clear()

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (ctrl *SessionController) getNewSession(c *gin.Context) app.Session {
	return app.NewRedisSession(c)
}

func (ctrl *SessionController) redirectToSignUp(c *gin.Context) {
	c.Redirect(http.StatusFound, "/users/new")
}
