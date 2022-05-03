package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"url_manager/app/forms"
	"url_manager/app/models"
	"url_manager/app/session"

	"url_manager/app/repositories"

	"github.com/gin-gonic/gin"
)

type URLCreateForm struct {
	Title string
	URL   string
}

type UrlsUri struct {
	UserId string `uri:"id"`
	UrlId  string `uri:"url_id"`
}

func (uri *UrlsUri) UserIdAsInt() int {
	value, err := strconv.Atoi(uri.UserId)
	if err != nil {
		panic(err)
	}
	return value
}

func (uri *UrlsUri) UrlIdAsInt() int {
	fmt.Println(uri.UrlId)
	value, err := strconv.Atoi(uri.UrlId)
	if err != nil {
		panic(err)
	}
	return value
}

func ShowURLs(c *gin.Context) {
	var uri UrlsUri
	c.ShouldBindUri(&uri)

	repo := repositories.NewUrlRepository()
	urls, err := repo.FindByUserID(uri.UserIdAsInt())
	if err != nil {
		return
	}

	session := session.NewRedisSession(c)

	c.HTML(
		http.StatusOK,
		"show_urls.html",
		gin.H{
			"login":   session.HasUserId(),
			"title":   "Show urls.",
			"user_id": uri.UserIdAsInt(),
			"urls":    urls,
		},
	)
}

func ShowURL(c *gin.Context) {
	session := session.NewRedisSession(c)
	c.HTML(
		http.StatusOK,
		"show_url.html",
		gin.H{
			"login": session.HasUserId(),
		},
	)
}

func NewURL(c *gin.Context) {
	var uri UrlsUri
	c.ShouldBindUri(&uri)
	session := session.NewRedisSession(c)

	c.HTML(
		http.StatusOK,
		"urls/new.html",
		gin.H{
			"login": session.HasUserId(),
			"id":    uri.UserIdAsInt(),
		},
	)
}

func CreateURL(c *gin.Context) {
	// リクエストからデータを抽出。
	var uri UrlsUri
	c.ShouldBindUri(&uri)

	var form forms.UrlCreateForm
	c.ShouldBind(&form)

	session := session.NewRedisSession(c)
	session.HasUserId()

	if !session.HasUserId() {
		log.Println("User id is not in the session.")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	urepo := repositories.NewUserRepository()
	user, err := urepo.FindById(session.GetUserId())
	if err != nil {
		log.Fatal("Can't find user. User id : ", session.GetUserId())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// メイン処理
	id := fmt.Sprintf("%d", user.ID)

	if uri.UserId != id {
		fmt.Println(id, uri.UserId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	url := models.NewUrl()
	url.Title = form.Title
	url.Url = "https://" + form.Url
	url.Description = form.Description
	url.Note = form.Note
	repo := repositories.NewUrlRepository()
	ownerId, _ := strconv.Atoi(id)
	repo.Create(ownerId, url)

	c.Redirect(http.StatusFound, "/users/"+uri.UserId+"/urls")
}

func EditURL(c *gin.Context) {
	session := session.NewRedisSession(c)
	c.HTML(
		http.StatusOK,
		"edit_url.html",
		gin.H{
			"login": session.HasUserId(),
		},
	)
}

func UpdateURL(c *gin.Context) {
	c.Redirect(http.StatusFound, "/urls")
}

func DeleteURL(c *gin.Context) {
	log.Println(c.Request.RequestURI)
	var uri UrlsUri
	c.ShouldBindUri(&uri)

	url := models.Url{}
	url.ID = uint(uri.UrlIdAsInt())
	fmt.Println(url.ID)

	repo := repositories.NewUrlRepository()
	err := repo.Delete(url)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
