package controller

import (
	"net/http"
	"strconv"
	"url_manager/model"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func IndexLink(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	links := repository.LinkIndex(loginId)

	c.HTML(http.StatusOK, "link/index.tmpl", gin.H{
		"IsLogin": isLogin,
		"LoginId": loginId,
		"Links":   links,
	})
}

func ShowLink(c *gin.Context) {

}

func NewLink(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	c.HTML(http.StatusOK, "link/new.tmpl", gin.H{
		"IsLogin": isLogin,
		"LoginId": loginId,
	})
}

func CreateLink(c *gin.Context) {
	loginId := c.GetString("loginId")
	var form LinkForm
	c.Bind(&form)

	repository.CreateLink(&model.Link{
		LoginId:     loginId,
		Title:       form.Title,
		Description: form.Description,
		URL:         form.URL,
	})

	c.Redirect(http.StatusFound, "/users/"+loginId+"/links")
}

func EditLink(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	linkId := c.Param("id")
	c.HTML(http.StatusOK, "link/edit.tmpl", gin.H{
		"LinkId":  linkId,
		"IsLogin": isLogin,
		"LoginId": loginId,
	})
}

func UpdateLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	loginId := c.GetString("loginId")
	var form LinkForm
	c.Bind(&form)

	repository.UpdateLink(&model.Link{
		Id:          id,
		Title:       form.Title,
		Description: form.Description,
		URL:         form.URL,
	})

	c.Redirect(http.StatusFound, "/users/"+loginId+"/links")
}

func DeleteLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	loginId := c.GetString("loginId")
	repository.DeleteLink(id, loginId)
	c.Redirect(http.StatusFound, "/users/"+loginId+"/links")
}
