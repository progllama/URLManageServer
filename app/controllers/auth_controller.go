package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	ErrorInvalidUserID   = errors.New("Invalid UserID")
	ErrorInvalidPassword = errors.New("Invalid Password")
	ErrorNotSignedIn     = errors.New("Not Logged In")
)

type AuthController struct{}

func (_ AuthController) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_up.tmpl", gin.H{})
}

func (_ AuthController) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "sing_in.tmpl", gin.H{})
}

func (_ AuthController) CreateSession(c *gin.Context) {
}

func (_ AuthController) DestroySession(c *gin.Context) {
	c.Redirect(302, "/users")
}

func (_ AuthController) SessionCheck(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("UserId")

	if uid == nil {
		c.Redirect(302, "/sing_in")
		c.Abort()
	} else {
		c.Set("UserId", uid)
		c.Next()
	}
}

func signIn(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

}

func signOut(c *gin.Context, UserId string) {
	session := sessions.Default(c)
	session.Set("UserId", UserId)
	session.Save()
}
