package controllers

import (
	"net/http"
	"url_manager/app/models"
	"url_manager/app/models/repositories"

	"github.com/gin-gonic/gin"
)

func ShowURLs(c *gin.Context) {
	var url models.URL
	c.BindJSON(&url)

	var r repositories.URLRepository

	_, err := r.Create(url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func ShowURL(c *gin.Context) {

}

func CreateURL(c *gin.Context) {

}

func UpdateURL(c *gin.Context) {

}

func DeleteURL(c *gin.Context) {

}
