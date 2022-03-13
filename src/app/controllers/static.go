package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{})
}
