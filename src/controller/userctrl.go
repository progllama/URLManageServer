package controllers

import (
	"log"
	"net/http"
	"strconv"
	"url_manager/database"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	// セッションチェック
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if loginId == nil {
		log.Fatal("not login")
	}
	strLoginId, ok := loginId.(string)
	if !ok {
		log.Fatal("cast error")
	}

	var user model.User
	result := database.DB.Where("login_id=?", strLoginId).First(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	result := database.DB.Where("id=?", id).First(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func UpdateUser(c *gin.Context) {
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if loginId == nil {
		log.Fatal("not login")
	}

	strLoginId, ok := loginId.(string)
	if !ok {
		log.Fatal("cast error")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var u model.User
	err = c.BindJSON(&u)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	result := database.DB.Where("id=?, login_id=?", id, strLoginId).First(&user)
	if result.Error != nil {
		log.Fatal(err)
	}

	result = database.DB.Model(&model.User{}).Where("id=?, login_id=?", user.ID, loginId).Update("name", u.Name)
	if result.Error != nil {
		log.Fatal(err)
	}

	c.Status(http.StatusOK)
}

func DeleteUser(ctx *gin.Context) {
	s := sessions.Default(ctx)
	loginId := s.Get("login_id")
	if loginId == nil {
		log.Fatal("not login")
	}

	strLoginId, ok := loginId.(string)
	if !ok {
		log.Fatal("cast error")
	}

	id, err := strconv.Atoi(ctx.Param(":id"))
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	result := database.DB.Where("id=?, login_id=?", id, strLoginId).First(&user)
	if result.Error != nil {
		log.Fatal(err)
	}

	result = database.DB.Where("id=?, login_id=?", id, strLoginId).Delete(&model.User{})
	if result.Error != nil {
		log.Fatal(err)
	}

	ctx.Status(http.StatusOK)
}
