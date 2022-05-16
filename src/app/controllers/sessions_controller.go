package controllers

import (
	"log"
	"net/http"
	"strconv"
	"url_manager/app/forms"
	"url_manager/app/repositories"
	"url_manager/app/session"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	userRepository repositories.UserRepository
}

func NewSessionController(userRepository repositories.UserRepository) *SessionController {
	return &SessionController{userRepository}
}

func (ctrl *SessionController) NewSession(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}

func (ctrl *SessionController) CreateSession(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// log.Println("Body: ", string(body))

	var form forms.LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := ctrl.userRepository.FindByLoginId(form.LoginId)
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

	c.Redirect(http.StatusFound, "/users/"+strconv.Itoa(int(user.ID)))
}

func (ctrl *SessionController) DestroySession(c *gin.Context) {
	session := ctrl.getNewSession(c)
	session.SetUserId(3)
	log.Println(session.ID())
	session.Clear()

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (ctrl *SessionController) getNewSession(c *gin.Context) session.Session {
	return session.NewMemSession(c)
}

func (ctrl *SessionController) redirectToSignUp(c *gin.Context) {
	c.Redirect(http.StatusFound, "/users/new")
}
