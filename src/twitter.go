package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var twitterScope = []string{
	"users.read",
	"tweet.read",
}

const (
	twitterRedirectURL = "http://127.0.0.1:8080/login"
	twitterProfileURL  = "https://api.twitter.com/2/users/me"
)

var TwitterConf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  twitterRedirectURL,
	Scopes:       twitterScope,
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://twitter.com/i/oauth2/authorize",
		TokenURL: "https://api.twitter.com/2/oauth2/token",
	},
}

type TwitterUser struct {
	ID string `json:"id"`
}

func (u *TwitterUser) LoginId() string {
	return u.ID
}

func TwitterOAuth2RedirectHandler(ctx *gin.Context) {
	ch := &OAuth2Handler{
		Context: ctx,
		Conf:    TwitterConf,
	}
	token, err := ch.Handle()
	if err != nil {
		return
	}

	// TwitterAPIでユーザデータの取得
	user, err := fetchTwitterUser(token)
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

func fetchTwitterUser(token *oauth2.Token) (*TwitterUser, error) {
	client := oauth2.NewClient(
		context.TODO(),
		oauth2.StaticTokenSource(token),
	)
	res, err := client.Get(twitterProfileURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var u TwitterUser
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func TwitterLoginURL(state string, codeChallenge string) string {
	return TwitterConf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
}
