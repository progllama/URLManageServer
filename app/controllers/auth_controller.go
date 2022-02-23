package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrorInvalidUserID   = errors.New("Invalid UserID")
	ErrorInvalidPassword = errors.New("Invalid Password")
	ErrorNotSignedIn     = errors.New("Not Logged In")
)

type AuthController struct{}

func (_ AuthController) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "sing_up.tmpl", gin.H{})
}

func (_ AuthController) SignOut(c *gin.Context) {

}

func (_ AuthController) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_up.tmpl", gin.H{})
}
