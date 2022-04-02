package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	session := sessions.Default(c)
	loggsin := session.Get(("uid")) != nil
	c.HTML(http.StatusOK, "static/about.html", gin.H{"title": "about", "logsin": loggsin, "id": session.Get(("uid"))})
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "static/contact.html", gin.H{"title": "contact"})
}
