package api

import (
	"url_manager/app/services"
	"url_manager/domain/models"

	"github.com/gin-gonic/gin"
)

// path parameter keys.
const (
	UserId = "userId"
)

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
	if err != nil {
		ctx.Error(err)
		return
	}

	response := api.service.Create(user)
	ctx.JSON(response.Code(), response.Body())
}

func (api *userApi) Update(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := api.service.Create(user)
	ctx.JSON(response.Code(), response.Body())
}

func (api *userApi) Delete(ctx *gin.Context) {
	userId := ctx.Param(UserId)

	response := api.service.Delete(userId)
	ctx.JSON(response.Code(), response.Body())
}
