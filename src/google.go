package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var googleScope = []string{
	"openid",
}

const (
	googleRedirectURL = "http://localhost:8080/login"
)

var GoogleConf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  googleRedirectURL,
	Scopes:       googleScope,
	Endpoint:     google.Endpoint,
}

type GoogleUser struct {
	ID string `json:"id"`
}

func (u *GoogleUser) LoginId() string {
	return u.ID
}

func GoogleLoginURL(state string, codeChallenge string) string {
	return GoogleConf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
}

func GoogleOAuth2RedirectHandler(ctx *gin.Context) {
	ch := &OAuth2Handler{
		Context: ctx,
		Conf:    GoogleConf,
	}
	token, err := ch.Handle()
	if err != nil {
		return
	}

	// GoogleAPIでユーザデータの取得
	user, err := fetchGoogleUser(token)
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
}

func fetchGoogleUser(token *oauth2.Token) (*GoogleUser, error) {
	service, err := goauth.NewService(
		context.TODO(),
		option.WithTokenSource(GoogleConf.TokenSource(context.TODO(), token)))
	if err != nil {
		return nil, err
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}
	return &GoogleUser{ID: userInfo.Id}, nil
}
