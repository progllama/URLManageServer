package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"url_manager/database"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"gorm.io/gorm"
)

var FacebookConf = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login",
	Scopes:      []string{"public_profile"},
	Endpoint:    facebook.Endpoint,
}

func FacebookAuth(c *gin.Context) {
	// stateのチェック
	s := sessions.Default(c)
	ss := s.Get("facebook state")
	qs := c.Query("state")
	if ss == nil || ss != qs {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// コードとトークンを交換
	code := c.Query("code")
	token, err := FacebookConf.Exchange(
		context.TODO(),
		code,
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// トークンとユーザ情報を交換
	client := oauth2.NewClient(
		context.TODO(),
		oauth2.StaticTokenSource(token),
	)
	res, err := client.Get("https://graph.facebook.com/v14.0/me")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var u struct {
		ID string `json:"id"`
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 外部アカウントと連携できてなければ連携
	var user model.User
	result := database.DB.Where("login_id", u.ID).First(&user)
	// レコードが見つからない以外場合のエラー、の時には終了
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// アカウントがない時
	if result.Error == gorm.ErrRecordNotFound {
		// アカウントの作成
		result = database.DB.Create(&model.User{LoginId: u.ID})
		if result.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// インデックスフォルダ作成のため新しく作成したアカウントのIDを取得
		var idChecker model.User
		result = database.DB.Where("login_id", u.ID).First(&idChecker)
		if result.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// インデックスフォルダの作成
		f := model.Folder{
			Name:   "index",
			UserID: idChecker.ID,
			Index:  true,
		}
		result := database.DB.Create(&f)
		if result.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// ユーザのインデックスIDを設定するためインデックスフォルダのIDを取得
		var fIdChecker model.Folder
		result = database.DB.Where("user_id=?, name=?", idChecker.ID, "index").First(&fIdChecker)
		if result.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// インデックスIDをユーザのIDインデックスに保存
		result = database.DB.Where("id=?", idChecker.ID).Update("index_id", fIdChecker.ID)
		if result.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	// セッションに保存
	s.Set("login_id", u.ID)
	s.Save()

	c.Status(http.StatusOK)
}

func GetFacebookLoginURL(c *gin.Context) {
	s := sessions.Default(c)
	state := randToken()
	s.Set("facebook state", state)
	s.Save()
	c.JSON(http.StatusOK, gin.H{"url": FacebookConf.AuthCodeURL(state)})
}
