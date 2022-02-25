package controllers

import (
	"net/http"
	"strconv"

	"url_manager/models/repository"

	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type UserController struct{}

// Index action: GET /users
func (_ UserController) Index(c *gin.Context) {
	var u repository.UserRepository
	_, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	}
}

// Create action: POST /users
func (_ UserController) Create(c *gin.Context) {
	var u repository.UserRepository
	p, err := u.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Redirect(302, "/users/"+strconv.Itoa(int(p.ID)))
	}
}

// Show action: Get /users/:id
func (_ UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	user, err := u.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.HTML(http.StatusOK, "user_show.tmpl", gin.H{"name": user.Name})
	}
}

// Update action: Put /users/:id
func (_ UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateByID(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (_ UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteByID(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + id + "Deleted"})
	return
}
