package auth

import (
	"context"
	"net/http"
	"url_manager/database"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

var GoogleConf = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login",
	Scopes:      []string{"public_profile"},
	Endpoint:    google.Endpoint,
}

func GoogleAuth(c *gin.Context) {
	s := sessions.Default(c)

	ss := s.Get("google state")
	qs := c.Query("state")

	if ss == nil || ss != qs {
		return
	}

	code := c.Query("code")
	token, err := GoogleConf.Exchange(
		context.TODO(),
		code,
	)
	if err != nil {
		return
	}

	service, err := goauth.NewService(
		context.TODO(),
		option.WithTokenSource(GoogleConf.TokenSource(context.TODO(), token)))
	if err != nil {
		return
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return
	}

	// 外部アカウントと連携できてなければ連携
	var user model.User
	result := database.DB.Where("login_id", userInfo.Id).First(&user)
	// レコードが見つからない以外場合のエラー、の時には終了
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// アカウントがない時
	if result.Error == gorm.ErrRecordNotFound {
		// アカウントの作成
		result = database.DB.Create(&model.User{LoginId: userInfo.Id})
		if result.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// インデックスフォルダ作成のため新しく作成したアカウントのIDを取得
		var idChecker model.User
		result = database.DB.Where("login_id", userInfo.Id).First(&idChecker)
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

	s.Set("login_id", userInfo.Id)
	s.Save()

	c.Status(http.StatusOK)
}

func GetGoogleLoginURL(c *gin.Context) {
	s := sessions.Default(c)
	state := randToken()
	s.Set("google state", state)
	s.Save()
	c.JSON(http.StatusOK, gin.H{"url": GoogleConf.AuthCodeURL(state)})
}
