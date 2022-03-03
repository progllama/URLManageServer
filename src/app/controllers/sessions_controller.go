package controllers

import (
	"net/http"
	"url_manager/app/models/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var loginParameter struct {
		Name     string
		Password string
	}
	c.BindJSON(loginParameter)

	session := sessions.Default(c)
	if session.Get("uid") == nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	u := repositories.UserRepository{}.GetByName(loginParameter.Name)
	if u == nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	if !authenticate(user.Name, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	session.Set("uid", u.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{})
}

func DestroySession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{})
}
