package controllers

import (
	"net/http"
	"strconv"
	"url_manager/app/models/repositories"

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

// // Update action: Put /users/:id
// func (_ UserController) Update(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var u repositories.UserRepository
// 	idInt, _ := strconv.Atoi(id)
// 	p, err := u.UpdateByID(idInt, c)

// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	} else {
// 		c.JSON(200, p)
// 	}
// }

// // Delete action: DELETE /users/:id
// func (_ UserController) Delete(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var u repositories.UserRepository
// 	idInt, _ := strconv.Atoi(id)
// 	if err := u.DeleteByID(idInt); err != nil {
// 		c.AbortWithStatus(403)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"success": "ID" + id + "Deleted"})
// 	return
// }
