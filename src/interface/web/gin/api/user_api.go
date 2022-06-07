package api

import (
	"url_manager/domain/models"
	"url_manager/domain/services"

	"github.com/gin-gonic/gin"
)

type UserApiConfig struct {
	UserIdPathParamKey string
}

func NewUserApiConfig() UserApiConfig {
	return UserApiConfig{}
}

type UserApi interface {
	Index(*gin.Context)
	Show(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userApi struct {
	service services.UserService
	config  UserApiConfig
}

func NewUserApi(s services.UserService, c UserApiConfig) UserApi {
	return &userApi{s, c}
}

func (api *userApi) Index(ctx *gin.Context) {
	response := api.service.FindUsers()
	api.assignResponse(ctx, response)
}

func (api *userApi) Show(ctx *gin.Context) {
	userId := api.extractUserId(ctx)
	response := api.service.FindUser(userId)
	api.assignResponse(ctx, response)
}

func (api *userApi) Create(ctx *gin.Context) {
	user, err := api.deserializeUser(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := api.service.Create(user)
	api.assignResponse(ctx, response)
}

func (api *userApi) Update(ctx *gin.Context) {
	user, err := api.deserializeUser(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := api.service.Create(user)
	api.assignResponse(ctx, response)
}

func (api *userApi) Delete(ctx *gin.Context) {
	userId := api.extractUserId(ctx)
	response := api.service.Delete(userId)
	api.assignResponse(ctx, response)
}

func (api *userApi) deserializeUser(ctx *gin.Context) (models.User, error) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (api *userApi) extractUserId(ctx *gin.Context) string {
	return ctx.Param(api.config.UserIdPathParamKey)
}

func (api *userApi) assignResponse(ctx *gin.Context, response services.UserServiceResponse) {
	ctx.JSON(response.Code, response.Body)
}
