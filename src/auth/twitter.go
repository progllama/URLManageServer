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
	"gorm.io/gorm"
)

var TwitterConf = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login",
	Scopes:      []string{"users.read", "tweet.read"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://twitter.com/i/oauth2/authorize",
		TokenURL: "https://api.twitter.com/2/oauth2/token",
	},
}

func TwitterAuth(c *gin.Context) {
	s := sessions.Default(c)

	ss := s.Get("twitter state")
	qs := c.Query("state")

	if ss == nil || ss != qs {
		return
	}

	code := c.Query("code")
	verifier := s.Get("verifier")
	if verifier == nil {
		return
	}
	token, err := TwitterConf.Exchange(
		context.TODO(),
		code,
		oauth2.SetAuthURLParam("code_verifier", verifier.(string)),
	)
	if err != nil {
		return
	}

	client := oauth2.NewClient(
		context.TODO(),
		oauth2.StaticTokenSource(token),
	)

	res, err := client.Get("https://api.twitter.com/2/users/me")
	if err != nil {
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var u struct {
		ID string `json:"id"`
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
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

	s.Set("login_id", u.ID)
	s.Save()

	c.Status(http.StatusOK)
}

func GetTwitterLoginURL(c *gin.Context) {

}

func TwitterLoginURL(c *gin.Context) {
	s := sessions.Default(c)
	state := randToken()
	s.Set("twitter state", state)
	s.Save()

	pkce := NewPKCE()
	s.Set("verifier", pkce.CodeVerifier())
	s.Save()

	c.JSON(http.StatusOK, gin.H{"url": TwitterConf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", pkce.CodeChallenge()),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)})
}
