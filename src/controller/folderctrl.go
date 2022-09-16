package controllers

import (
	"log"
	"net/http"
	"url_manager/database"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexFolder(c *gin.Context) {
	userId := c.Param("user_id")

	s := sessions.Default(c)
	loginId := s.Get("login_id")

	if userId != loginId {
		log.Fatal("please login")
	}

	var index model.Folder
	database.DB.Where("user_id=?, index=?", userId, true).First(&index)

	var folders []model.Folder
	result := database.DB.Select("id, name").Where("user_id=?, parent=?", userId, index.ID).Find(&folders)
	if result.Error != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, folders)
}

func GetFolders(c *gin.Context) {
	userId := c.Param("user_id")
	s := sessions.Default(c)
	loginId := s.Get("login_id")

	if userId != loginId {
		log.Fatal("please login")
	}

	folderId := c.Param("folder_id")
	var folder model.Folder
	result := database.DB.Select("id, name").Where("user_id=?, id=?", userId, folderId).First(&folder)
	if result.Error != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	var children model.Folder
	result = database.DB.Select("id, name").Where("user_id=?, parent=?", userId, folder.ID).Find(&children)
	if result.Error != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{"folder": folder, "children": children})
}

func UpdateFolder(c *gin.Context) {
	userId := c.Param("user_id")
	s := sessions.Default(c)
	loginId := s.Get("login_id")

	if userId != loginId {
		log.Fatal("please login")
	}

	var req model.User
	c.BindJSON(req)

	folderId := c.Param("folder_id")
	result := database.DB.Where("user_id=?, id=?", userId, folderId).Update("name", req.Name)
	if result.Error != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func DeleteFolder(c *gin.Context) {
	userId := c.Param("user_id")
	s := sessions.Default(c)
	loginId := s.Get("login_id")

	if userId != loginId {
		log.Fatal("please login")
	}

	folderId := c.Param("folder_id")
	result := database.DB.Where("id=?", folderId).Delete(&model.User{})
	if result.Error != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}
