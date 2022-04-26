package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"url_manager/app"
	"url_manager/app/forms"
	"url_manager/app/models"
	"url_manager/app/repositories"
	"url_manager/app/uris"

	"github.com/gin-gonic/gin"
)

var (
	ErrCantExtractUserId = errors.New("can't extract user id.")
)

type UsersController struct {
	repo *repositories.UserRepository
}

func NewUserController() *UsersController {
	return &UsersController{
		repo: repositories.NewUserRepository(),
	}
}

func (ctrl *UsersController) ShowAll(c *gin.Context) {
	users, err := ctrl.repo.AllIdAndNames()
	if err != nil {
		c.Error(err)
		return
	}

	session := app.NewRedisSession(c)
	login := session.HasUserId()

	c.HTML(
		http.StatusOK,
		"show_users.html",
		gin.H{
			"login": login,
			"users": users,
		},
	)
}

func (ctrl *UsersController) Show(ctx *gin.Context) {
	userId, err := ctrl.extractUserId(ctx)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := ctrl.repo.FindById(userId)
	if err != nil {
		log.Fatal(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	session := app.NewRedisSession(ctx)
	login := session.HasUserId()

	ctx.HTML(
		http.StatusOK,
		"show_user.html",
		gin.H{
			"login": login,
			"id":    user.ID,
			"user":  user,
		},
	)
}

func (ctrl *UsersController) extractUserId(ctx *gin.Context) (int, error) {
	session := app.NewRedisSession(ctx)
	if session.HasUserId() {
		return session.GetUserId(), nil
	}

	var uri uris.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err == nil {
		return uri.GetUserId(), nil
	}

	return 0, ErrCantExtractUserId
}

func (ctrl *UsersController) New(c *gin.Context) {
	session := app.NewRedisSession(c)

	c.HTML(
		http.StatusOK,
		"new_user.html",
		gin.H{
			"login": session.HasUserId(),
			"title": "NewUser",
		},
	)
}

func (ctrl *UsersController) Create(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
	var form forms.UserCreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.Error(err)
		return
	}

	log.Println("Success form binding.")

	exist, err := ctrl.repo.Exists(models.User{Name: form.Name})
	if err != nil {
		log.Fatal(err)
		return
	}
	if exist {
		log.Println("Name is not unique.")
		return
	}

	log.Println("Name is unique.")

	err = ctrl.repo.Create(form.Name, form.LoginId, form.Password)
	if err != nil {
		c.Error(err)
		return
	}

	log.Println("Success create.")

	c.Redirect(http.StatusFound, "/login")
}

func (ctrl *UsersController) Edit(c *gin.Context) {
	session := app.NewRedisSession(c)
	c.HTML(http.StatusOK, "edit_user.html", gin.H{"login": session.HasUserId()})
}

func (ctrl *UsersController) Update(c *gin.Context) {
	var uri uris.UserUri
	c.ShouldBindUri(&uri)

	var form forms.UserEditForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.Error(err)
		return
	}

	ctrl.repo.Update(uri.ToInt(), form.Name, form.LoginId, form.Password)
	if err != nil {
		c.Error(err)
	}

	c.Redirect(http.StatusFound, "/users/"+strconv.Itoa(uri.ToInt()))
}

func (ctrl *UsersController) Delete(c *gin.Context) {
	var uri uris.UserUri
	c.ShouldBind(uri)

	if err := ctrl.repo.Delete(uri.ToInt()); err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/home")
}
