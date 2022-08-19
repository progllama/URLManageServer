package controllers

import (
	"net/http"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	s := sessions.Default(ctx)
	id := s.Get("id")
	if id != nil && id != "" && id == ctx.Param("user_id") {
		return
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, emptyBody)
}

func Login(ctx *gin.Context) {
	s := sessions.Default(ctx)
	if ok, u := authenticateUser(ctx); ok {
		s.Set("id", u.ID)
		s.Save()

		ctx.JSON(http.StatusOK, u)
		return
	}

	ctx.JSON(http.StatusNotFound, emptyBody)
}

func Logout(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	ctx.JSON(http.StatusOK, emptyBody)
}

func authenticateUser(ctx *gin.Context) (bool, *model.User) {
	var credential model.User
	ctx.BindJSON(&credential)

	repo := getUserRepo()
	u := repo.GetByUserId(credential.UserID)
	if u.Authenticate(&credential) {
		return true, &u
	} else {
		return false, nil
	}
}
