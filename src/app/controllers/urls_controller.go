package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"url_manager/app/forms"
	"url_manager/app/models"

	"url_manager/app/repositories"

	"github.com/gin-contrib/sessions"
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
	value, err := strconv.Atoi(uri.UserId)
	if err != nil {
		panic(err)
	}
	return value
}

func ShowURLs(c *gin.Context) {
	var uri UrlsUri
	c.ShouldBindUri(&uri)

	repo := repositories.NewUrlRepository()
	urls, err := repo.FindByUserID(uri.UrlIdAsInt())
	fmt.Println(urls[2].Note)
	if err != nil {
		return
	}

	c.HTML(http.StatusOK, "show_urls.html", gin.H{"title": "Show urls.", "user_id": uri.UserIdAsInt(), "urls": urls})
}

func ShowURL(c *gin.Context) {
	c.HTML(http.StatusOK, "show_url.html", gin.H{})
}

func NewURL(c *gin.Context) {
	var uri UrlsUri
	c.ShouldBindUri(&uri)

	c.HTML(http.StatusOK, "urls/new.html", gin.H{"id": uri.UserIdAsInt()})
}

func CreateURL(c *gin.Context) {
	// リクエストからデータを抽出。
	var uri UrlsUri
	c.ShouldBindUri(&uri)

	var form forms.UrlCreateForm
	c.ShouldBind(&form)

	session := sessions.Default(c)
	loginId := session.Get("login_id")
	if loginId == nil {
		c.Abort()
		return
	}

	urepo := repositories.NewUserRepository()
	user, err := urepo.FindByLoginId(loginId.(string))
	if err != nil {
		c.Abort()
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
	c.HTML(http.StatusOK, "edit_url.html", gin.H{})
}

func UpdateURL(c *gin.Context) {
	c.Redirect(http.StatusFound, "/urls")
}

func DeleteURL(c *gin.Context) {
	url := models.Url{}
	id, _ := strconv.Atoi(c.Param("urlID"))
	url.ID = uint(id)

	repo := repositories.PostgreSQLURLRepository{}
	repo.Delete(url)
	fmt.Println(url.ID)

	session := sessions.Default(c)
	userID := fmt.Sprintf("%d", session.Get("uid"))

	c.Redirect(302, "/users/"+userID)
}
