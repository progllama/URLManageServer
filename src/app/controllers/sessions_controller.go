package controllers

import (
	"net/http"
	"url_manager/app/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	urepo repositories.IUserRepository
}

func (ctrl *UserController) CreateSession(ctx *gin.Context) {
	var loginParameter struct {
		Name     string
		Password string
	}
	err := ctx.BindJSON(&loginParameter)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(ctx)
	if IsLoggedIn(session) {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	u, err := repositories.UserRepository{}.GetByName(loginParameter.Name)
	if Authenticate(u.Name, u.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	Login(session, u.ID)

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})
}

func IsLoggedIn(s sessions.Session) bool {
	return s.Get("uid") != nil
}

func Authenticate(name string, password string) bool {
	u, err := repositories.UserRepository{}.GetByName(name)
	if err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func Login(s sessions.Session, id uint) {
	s.Set("uid", id)
}

func DestroySession(c *gin.Context) {
	session := sessions.Default(c)
	Logout(session)
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(s sessions.Session) {
	s.Clear()
	s.Save()
}
