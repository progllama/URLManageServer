package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"url_manager/app/models"
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

	urlrepo := repositories.DefaultURLRepositoryImpl{}
	urls, err := urlrepo.GetByUserID(uid)

	safeUrls := make([]SafeURL, len(urls))

	for i, v := range urls {
		safeUrls[i] = SafeURL{
			fmt.Sprintf("%d", v.ID),
			v.Title,
			v.URL,
		}
	}

	fmt.Println(c.GetBool("logsin"))

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.HTML(http.StatusOK, "user/show.html", gin.H{
			"title":  "ShowUser",
			"name":   user.Name,
			"id":     user.ID,
			"urls":   safeUrls,
			"logsin": c.GetBool("logsin"),
		})
	}
}

type SafeURL struct {
	ID    string
	Title string
	Url   string
}

func NewUser(c *gin.Context) {
	c.HTML(http.StatusOK, "user/new.html", gin.H{"title": "NewUser"})
}

func CreateUser(c *gin.Context) {
	c.Request.ParseForm()
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")

	r := repositories.UserRepository{}
	_, err := r.Exists(name)
	if err != nil {
		fmt.Println(err)
		c.Redirect(302, "/users/new")
		c.Abort()
		return
	}

	err = r.Create(models.User{
		Name:     name,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
		c.Redirect(302, "/users")
		c.Abort()
		return
	}

	u, err := r.GetByName(name)

	sessoin := sessions.Default(c)
	Login(sessoin, u.ID)

	c.Redirect(302, "/login")
	c.Abort()
}

func EditUser(c *gin.Context) {

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
