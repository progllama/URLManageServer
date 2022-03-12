package controllers

import (
	"net/http"
	"strconv"
	"url_manager/app/models"
	"url_manager/app/models/repositories"

	"github.com/gin-gonic/gin"
)

func ShowURLs(c *gin.Context) {
	var url models.URL
	c.BindJSON(&url)

	r := repositories.DefaultURLRepositoryImpl{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
	}
	urls, err := r.GetByUserID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}

func ShowURL(c *gin.Context) {
	var url models.URL
	c.BindJSON(&url)

	var r repositories.URLRepository

	err := r.Create(url)
}

func CreateURL(c *gin.Context) {

}

func UpdateURL(c *gin.Context) {

}

func DeleteURL(c *gin.Context) {

}
