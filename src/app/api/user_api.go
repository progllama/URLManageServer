package api

import (
	"net/http"
	"url_manager/app/models"
	"url_manager/app/repositories"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	repo repositories.UserRepository
}

func NewUserApi(r repositories.UserRepository) *UserApi {
	return &UserApi{r}
}

func (api *UserApi) Index(ctx *gin.Context) {
	users, err := api.repo.All()
	if err != nil {
		ctx.JSON(http.StatusNotFound, []models.User{})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (api *UserApi) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *UserApi) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
