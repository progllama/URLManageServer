package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"url_manager/app/models"
	"url_manager/app/models/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowURLs(c *gin.Context) {

}

func ShowURL(c *gin.Context) {

}

func NewURL(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("uid")

	c.HTML(http.StatusOK, "urls/new.html", gin.H{"id": id})
}

func CreateURL(c *gin.Context) {
	session := sessions.Default(c)
	id := fmt.Sprintf("%d", session.Get("uid"))

	c.Request.ParseForm()
	title := c.Request.FormValue("title")
	url := c.Request.FormValue("url")

	intid, _ := strconv.Atoi(id)

	urlModel := models.URL{
		Title:  title,
		URL:    url,
		UserID: intid,
	}

	r := repositories.DefaultURLRepositoryImpl{}
	r.Create(urlModel)
	a, _ := r.GetAll()
	fmt.Println(a)
	c.Redirect(302, "/users/"+id)
}

func EditURL(c *gin.Context) {

}

func UpdateURL(c *gin.Context) {

}

func DeleteURL(c *gin.Context) {
	url := models.URL{}
	id, _ := strconv.Atoi(c.Param("urlID"))
	url.ID = uint(id)

	repo := repositories.DefaultURLRepositoryImpl{}
	repo.Destroy(url)
	fmt.Println(url.ID)

	session := sessions.Default(c)
	userID := fmt.Sprintf("%d", session.Get("uid"))

	c.Redirect(302, "/users/"+userID)
}
