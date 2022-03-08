package controllers

import (
	"net/http"
	"url_manager/app/models/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateSession(c *gin.Context) {
	var loginParameter struct {
		Name     string
		Password string
	}
	err := c.BindJSON(&loginParameter)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	if IsLoggedIn(session) {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	u, err := repositories.UserRepository{}.GetByName(loginParameter.Name)
	if Authenticate(u.Name, u.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	Login(session, u.ID)

	c.JSON(http.StatusOK, gin.H{"result": "success"})
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
