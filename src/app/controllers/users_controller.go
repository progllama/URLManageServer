package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"url_manager/app/models"
	"url_manager/app/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserCreatForm struct {
	Name     string `form:"name"`
	LoginId  string `form:"login_id"`
	Password string `form:"password"`
}

type UserURI struct {
	Id string `uri:"id"`
}

func (uri *UserURI) ToInt() int {
	id, err := strconv.Atoi(uri.Id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

type UserEditForm struct {
	Name     string `form:"name"`
	LoginId  string `form:"login_id"`
	Password string `form:"password"`
}

type UsersController struct {
	repo *repositories.UserRepository
}

func NewUserController() *UsersController {
	return &UsersController{
		repo: repositories.NewUserRepository(),
	}
}

func (ctrl *UsersController) ShowAll(c *gin.Context) {
	users, err := ctrl.repo.All()
	if err != nil {
		c.Error(err)
		return
	}

	c.HTML(http.StatusOK, "show_users.html", gin.H{"loggedin": ctrl.logsin(c), "users": users})
}

func (ctrl *UsersController) Show(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
	var uri UserURI
	err := c.ShouldBindUri(&uri)

	log.Println("Success form binding.")
	log.Println(uri.Id)

	user, err := ctrl.repo.FindByID(uri.ToInt())
	if err != nil {
		c.Error(err)
		return
	}

	log.Println("Success find user.")

	c.HTML(http.StatusOK, "show_user.html", gin.H{"loggedin": ctrl.logsin(c), "id": user.ID, "user": user})
}

func (ctrl *UsersController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new_user.html", gin.H{"loggedin": ctrl.logsin(c), "title": "NewUser"})
}

func (ctrl *UsersController) Create(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
	var form UserCreatForm
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
	c.HTML(http.StatusOK, "edit_user.html", gin.H{"loggedin": ctrl.logsin(c)})
}

func (ctrl *UsersController) Update(c *gin.Context) {
	var uri UserURI
	c.ShouldBindUri(&uri)

	var form UserEditForm
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
	var uri UserURI
	c.ShouldBind(uri)

	if err := ctrl.repo.Delete(uri.ToInt()); err != nil {
		return
	}

	c.Redirect(http.StatusFound, "/home")
}

func (ctrl *UsersController) logsin(c *gin.Context) bool {
	sessoin := sessions.Default(c)
	id := sessoin.Get("login_id")
	fmt.Println(id)
	return id != nil
}
