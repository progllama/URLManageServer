package controller

import (
	"net/http"
	"strconv"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func IndexFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	folders := repository.IndexFolder(loginId)
	c.HTML(http.StatusOK, "folder/index.tmpl", gin.H{
		"LoginId": loginId,
		"IsLogin": isLogin,
		"Folders": folders,
	})
}

func ShowFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	c.HTML(http.StatusOK, "folder/show.tmpl", gin.H{
		"LoginId": loginId,
		"IsLogin": isLogin,
	})
}

func NewFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	c.HTML(http.StatusOK, "folder/new.tmpl", gin.H{
		"LoginId": loginId,
		"IsLogin": isLogin,
	})
}

func CreateFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	var form FolderForm
	c.Bind(&form)
	repository.CreateFolder(loginId, form.Title)
	c.Redirect(http.StatusFound, "/users/"+loginId+"/folders")
}

func EditFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	isLogin := c.GetBool("isLogin")
	id, _ := strconv.Atoi(c.Param("id"))
	c.HTML(http.StatusOK, "folder/edit.tmpl", gin.H{
		"LoginId": loginId,
		"IsLogin": isLogin,
		"Id":      id,
	})
}

func UpdateFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	id, _ := strconv.Atoi(c.Param("id"))
	var form FolderForm
	c.Bind(&form)
	repository.UpdateFolder(id, form.Title)
	c.Redirect(http.StatusFound, "/users/"+loginId+"/folders")
}

func DeleteFolder(c *gin.Context) {
	loginId := c.GetString("loginId")
	id, _ := strconv.Atoi(c.Param("id"))
	repository.DeleteFolder(loginId, id)
	c.Redirect(http.StatusFound, "/users/"+loginId+"/folders")
}
