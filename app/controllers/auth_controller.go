package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorInvalidUserID   = errors.New("Invalid UserID")
	ErrorInvalidPassword = errors.New("Invalid Password")
	ErrorNotSignedIn     = errors.New("Not Logged In")
)

type AuthController struct{}

func (_ AuthController) SignIn(c *gin.Context) {

}

func (_ AuthController) SignOut(c *gin.Context) {

}
