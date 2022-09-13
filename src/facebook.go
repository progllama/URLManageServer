package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var facebookScope = []string{
	"public_profile",
}

const (
	facebookRedirectURL = "http://localhost:8080/login"
	facebookProfileURL  = "https://graph.facebook.com/v14.0/me"
)

var FacebookConf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  facebookRedirectURL,
	Scopes:       facebookScope,
	Endpoint:     facebook.Endpoint,
}

type FacebookUser struct {
	ID string `json:"id"`
}

func (u *FacebookUser) LoginId() string {
	return u.ID
}

func FacebookLoginURL(state string) string {
	return FacebookConf.AuthCodeURL(state)
}

func FaceBookOAuth2RedirectHandler(ctx *gin.Context) {
	ch := &OAuth2Handler{
		Context: ctx,
		Conf:    GoogleConf,
	}
	token, err := ch.Handle()
	if err != nil {
		return
	}

	// FacebookAPIでユーザデータの取得
	user, err := fetchFacebookUser(token)
	if err != nil {
		ctx.JSON(
			ErrorCode(err),
			gin.H{errorKey: err.Error()},
		)
		return
	}

	// ユーザデータがアプリに登録されていなければユーザの新規作成
	account, err := createUserIfNotExist(user)
	if err != nil {
		ctx.JSON(
			ErrorCode(err),
			gin.H{errorKey: ErrAccessDatabase.Error() + err.Error()},
		)
		return
	}

	// ログイン(セッションへの保存)
	err = login(ctx, account)
	if err != nil {
		ctx.JSON(
			ErrorCode(err),
			gin.H{errorKey: ErrOnExchange.Error() + err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{accountIdKey: account.ID})
}

func fetchFacebookUser(token *oauth2.Token) (*FacebookUser, error) {
	client := oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(token))
	res, err := client.Get(facebookProfileURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var u FacebookUser
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
