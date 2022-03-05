package controllers

import (
	"net/http"
	"strconv"
	"url_manager/app/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	var repo repositories.UserRepository
	users, err := repo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func ShowUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if uint(id) != sessions.Default(c).Get("uid") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong id or not logged in."})
		return
	}

	var u repositories.UserRepository
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	user, err := u.GetByID(uid)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"name": user.Name, "id": user.ID})
	}
}

func CreateUser(c *gin.Context) {
	var repo repositories.UserRepository
	// TODO gin.Contextへの依存がレポジトリにも発生しているのでこれをよしとしていくか
	//      依存しない方向でいくか。
	u, err := repo.CreateModel(c)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		// TODO session コントローラのログイン処理とまとめる
		s := sessions.Default(c)
		s.Set("uid", u.ID)
		s.Save()
	}
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if uint(id) != sessions.Default(c).Get("uid") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong id or not logged in."})
		return
	}

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	user, err := repositories.UserRepository{}.UpdateByID(uid, c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"name": user.Name, "id": user.ID})
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var u repositories.UserRepository
	if err := u.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + strconv.Itoa(id) + "Deleted"})
	return
}
