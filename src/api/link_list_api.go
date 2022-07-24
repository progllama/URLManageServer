package api

import (
	"net/http"
	"strconv"
	"url_manager/models"
	"url_manager/repositories"

	"github.com/gin-gonic/gin"
)

type LinkListsApi struct {
	linkListRepository repositories.LinkListRepository
}

func NewLinkListsApi(llrepo repositories.LinkListRepository) *LinkListsApi {
	return &LinkListsApi{llrepo}
}

func (api *LinkListsApi) Create(ctx *gin.Context) {
	var list models.LinkList
	ctx.BindJSON(&list)

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	list.UserID = userID

	api.linkListRepository.Add(list)
	lists, err := api.linkListRepository.FindByUserId(userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	cnt := 0
	for i, v := range lists {
		if v.Title == list.Title {
			cnt = i
			break
		}
	}

	list.ID = lists[cnt].ID
	ctx.JSON(http.StatusOK, list)
}

func (api *LinkListsApi) Delete(ctx *gin.Context) {
	listId, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = api.linkListRepository.Remove(listId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
