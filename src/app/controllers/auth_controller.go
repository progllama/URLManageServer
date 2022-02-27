package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (_ AuthController) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_up.tmpl", gin.H{})
}

func (_ AuthController) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_in.tmpl", gin.H{})
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
