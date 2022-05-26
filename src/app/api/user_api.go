package api

import (
	"url_manager/app/models"
	"url_manager/app/services"

	"github.com/gin-gonic/gin"
)

// This responsibilities
// 1. extract user data.
// 2. call service with user data.
// 3. set response with data returned by service.
type UserApi interface {
	Index()
	Show()
	Create()
	Update()
	Delete()
}

type userApi struct {
	service services.UserService
}

func NewUserApi(s services.UserService) *userApi {
	return &userApi{s}
}

func (api *userApi) Index(ctx *gin.Context) {
	response := api.service.FindUsers()
	ctx.JSON(response.Code(), response.Body())
}

func (api *userApi) Show(ctx *gin.Context) {
	userId := ctx.Param("userId")

	response := api.service.FindUser(userId)
	ctx.JSON(response.Code(), response.Body())
}

func (api *userApi) Create(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	response := api.service.Create(user)
	ctx.JSON(response.Code(), response.Body())
}

func (api *userApi) Update(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	response := api.service.Create(user)
	ctx.JSON(response.Code(), response.Body())
}

func (api *userApi) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")

	response := api.service.Delete(userId)
	ctx.JSON(response.Code(), response.Body())
}
