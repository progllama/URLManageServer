package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"url_manager/models"
	"url_manager/repositories"

	"github.com/gin-gonic/gin"
)

type LinksApi interface {
	Index(*gin.Context)
	Show(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type linksApi struct {
	userRepository             repositories.UserRepository
	linkRepository             repositories.LinkRepository
	linkListRepository         repositories.LinkListRepository
	linkListRelationRepository repositories.LinkListRelationRepository
}

func NewLinksApi(
	userRepository repositories.UserRepository,
	linkRepository repositories.LinkRepository,
	linkListRepository repositories.LinkListRepository,
	linkListRelationRepository repositories.LinkListRelationRepository,
) LinksApi {
	return &linksApi{userRepository, linkRepository, linkListRepository, linkListRelationRepository}
}

func (api *linksApi) Index(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)

	user, err := api.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userIdInParam, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if int(user.ID) != userIdInParam {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	lists, err := api.linkListRepository.FindByUserId(int(user.ID))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	relations, err := api.linkListRelationRepository.FindByUserId(int(user.ID))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var roots []int
	for _, list := range lists {
		isChildList := false
		for _, relation := range relations {
			if int(list.ID) == relation.ChildListID {
				isChildList = true
				break
			}
		}

		if !isChildList {
			roots = append(roots, int(list.ID))
		}
	}

	var urlsToEachList []string
	for _, id := range roots {
		url := "/users/" + fmt.Sprint(user.ID) + "/links/" + fmt.Sprint(id)
		urlsToEachList = append(urlsToEachList, url)
	}

	urlsJson, err := json.Marshal(urlsToEachList)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, urlsJson)
}

func (api *linksApi) Show(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)

	user, err := api.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userIdInParam, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if int(user.ID) != userIdInParam {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	listId, err := strconv.Atoi(ctx.Param("link_list_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	relations, err := api.linkListRelationRepository.FindByParentId(listId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var linkData []string
	for _, relation := range relations {
		linkData = append(linkData, fmt.Sprint(relation.ChildListID))
	}

	links, err := api.linkRepository.FindByListId(listId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, link := range links {
		linkData = append(linkData, link.Url)
	}

	ctx.JSON(http.StatusOK, linkData)
}

func (api *linksApi) Create(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)

	user, err := api.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userIdInParam, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if int(user.ID) != userIdInParam {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var link models.Link
	ctx.Bind(&link)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if link.UserID != int(user.ID) {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = api.linkRepository.Add(link)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *linksApi) Update(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)

	user, err := api.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userIdInParam, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if int(user.ID) != userIdInParam {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var link models.Link
	ctx.Bind(&link)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if link.UserID != int(user.ID) {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = api.linkRepository.Update(link)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (api *linksApi) Delete(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)

	user, err := api.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userIdInParam, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if int(user.ID) != userIdInParam {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var link models.Link
	ctx.Bind(&link)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if link.UserID != int(user.ID) {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = api.linkRepository.Remove(int(link.ID))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
