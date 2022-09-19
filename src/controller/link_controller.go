package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CreateLink(c *gin.Context) {
	c.Status(http.StatusOK)
}
