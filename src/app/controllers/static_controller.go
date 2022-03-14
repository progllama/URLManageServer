package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "static/about.html", gin.H{"title": "about"})
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "static/contact.html", gin.H{"title": "contact"})
}
