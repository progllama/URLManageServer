package controllers

import (
	"net/http"
	"url_manager/app/models/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// TODO
// ログインできるか確認。 canLogin
// 本人確認 authenticate
// ログイン login
// ログアウト logout

func CreateSession(c *gin.Context) {
	var loginParameter struct {
		Name     string
		Password string
	}
	err := c.BindJSON(&loginParameter)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	session := sessions.Default(c)
	if session.Get("uid") != nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	u, err := repositories.UserRepository{}.GetByName(loginParameter.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginParameter.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	session.Set("uid", u.ID)

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func DestroySession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{})
}
