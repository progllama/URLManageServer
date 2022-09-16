package controllers

import (
	"net/http"
	"strconv"
	"url_manager/database"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexLink(c *gin.Context) {
	// ユーザ取得
	userId := c.Param("user_id")
	var user model.User
	result := database.DB.Select("id, login_id").Where("id=?", userId).First(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// セッションチェック
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if user.LoginId != loginId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ユーザIDとfolderIdが一致するリンクを取得する
	folderId := c.Param("folder_id")
	var links []model.Link
	result = database.DB.Select("title, url").Where("user_id=?, folder_id=?", userId, folderId).Find(&links)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, links)
}

func GetLink(c *gin.Context) {
	// ユーザ取得
	userId := c.Param("user_id")
	var user model.User
	result := database.DB.Select("id, login_id").Where("id=?", userId).First(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// セッションチェック
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if user.LoginId != loginId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ユーザIDとフォルダIDとリンクIDが一致するリンクを取得する
	folderId := c.Param("folder_id")
	linkId := c.Param("link_id")
	var link model.Link
	result = database.DB.Select("title, url").
		Where("id=?, user_id=?, folder_id=?", linkId, userId, folderId).
		Find(&link)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, link)
}

func CreateLink(c *gin.Context) {
	// ユーザ取得
	userId := c.Param("user_id")
	var user model.User
	result := database.DB.Select("id, login_id").Where("id=?", userId).First(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// セッションチェック
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if user.LoginId != loginId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// リクエストのでシリアライズ
	var link model.Link
	err := c.BindJSON(&link)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// バリデーション
	valid := true
	valid = valid && link.Title != ""
	valid = valid && link.URL != ""
	// ユーザIDのチェック
	userIdint, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	valid = valid && link.UserID == uint(userIdint)
	if !valid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// フォルダIDのチェック
	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	valid = valid && link.FolderID == uint(folderId)
	if !valid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 作成
	result = database.DB.Create(&link)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, link)
}

func UpdateLink(c *gin.Context) {
	// ユーザ取得
	userId := c.Param("user_id")
	var user model.User
	result := database.DB.Select("id, login_id").Where("id=?", userId).First(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// セッションチェック
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if user.LoginId != loginId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// リクエストのでシリアライズ
	var link model.Link
	err := c.BindJSON(&link)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// バリデーション
	valid := true
	valid = valid && link.Title != ""
	valid = valid && link.URL != ""
	// ユーザIDのチェック
	userIdint, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	valid = valid && link.UserID == uint(userIdint)
	if !valid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// フォルダIDのチェック
	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	valid = valid && link.FolderID == uint(folderId)
	if !valid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 更新
	result = database.DB.Where("link_id=?", c.Param("link_id")).Updates(
		model.Link{
			Title: link.Title,
			URL:   link.URL,
		},
	)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, link)
}

func DeleteLink(c *gin.Context) {
	// ユーザ取得
	userId := c.Param("user_id")
	var user model.User
	result := database.DB.Select("id, login_id").Where("id=?", userId).First(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// セッションチェック
	s := sessions.Default(c)
	loginId := s.Get("login_id")
	if user.LoginId != loginId {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ユーザIDとフォルダIDとリンクIDが一致するリンクを削除する
	folderId := c.Param("folder_id")
	linkId := c.Param("link_id")
	link := model.Link{}
	linkIdint, err := strconv.Atoi(linkId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	link.ID = uint(linkIdint)
	result = database.DB.Select("title, url").
		Where("id=?, user_id=?, folder_id=?", linkId, userId, folderId).
		Delete(&link)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, link)
}
