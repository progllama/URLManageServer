package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	google "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

// please initialize token yourself
var conf oauth2.Config

// 処理に成功した場合にはコード200(OK)、{"accountId": id}を返却
// 失敗した場合にはコード500(InternalServerError)を返却
func LoginGoogle(ctx *gin.Context) {
	// ログイン済みであればログインしているユーザのIDを返却
	if isLogin(ctx) {
		id, err := getSessionAccountID(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"accountId": id})
		return
	}

	// 認証情報(OpenID)の取得
	openID, err := getOpenID(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 認証情報の識別(アカウントが存在するか確認)し,なければアカウントの作成
	accountExists, err := identify(openID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !accountExists {
		err = createAccount(openID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	// アカウントのIDを取得
	id, err := getAccountID(openID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// アカウントのIDをセッションに保存
	err = saveSession(ctx, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accountId": id})
}

// ログイン済みか確認
func isLogin(ctx *gin.Context) bool {
	s := sessions.Default(ctx)
	id := s.Get("id")
	return id != nil
}

// セッションのアカウントIDを取得
func getSessionAccountID(ctx *gin.Context) (int, error) {
	s := sessions.Default(ctx)
	id, ok := s.Get("id").(int)
	if ok {
		return id, nil
	}
	return 0, errors.New("can't cast int")
}

// openIDを取得
func getOpenID(ctx *gin.Context) (string, error) {
	// トークンを取得
	token, err := getToken(ctx)
	if err != nil {
		return "", err
	}
	// トークンと情報の交換
	openId, err := exchangeToken(ctx, token)
	if err != nil {
		return "", err
	}
	return openId, nil
}

func getToken(ctx *gin.Context) (*oauth2.Token, error) {
	code, err := getCode(ctx)
	if err != nil {
		return nil, err
	}
	token, err := exchangeCode(code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func exchangeCode(code string) (*oauth2.Token, error) {
	return conf.Exchange(context.TODO(), code)
}

func getCode(ctx *gin.Context) (string, error) {
	// ステートの比較
	if !isStateValid(ctx) {
		return "", errors.New("invalid state")
	}
	return ctx.Query("code"), nil
}

func isStateValid(ctx *gin.Context) bool {
	s := sessions.Default(ctx)
	state := s.Get("state")
	return state == ctx.Query("state")
}

func exchangeToken(ctx *gin.Context, token *oauth2.Token) (string, error) {
	service, err := google.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, token)))
	if err != nil {
		return "", err
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return "", err
	}
	return userInfo.Id, nil
}

func identify(openId string) (bool, error) {
	repo := AccountRepository{}
	return repo.Exists(openId)
}

func createAccount(openId string) error {
	repo := AccountRepository{}
	return repo.Create(openId)
}

func getAccountID(openId string) (int, error) {
	repo := AccountRepository{}
	return repo.Find(openId)
}

func saveSession(ctx *gin.Context, id int) error {
	s := sessions.Default(ctx)
	s.Set("id", id)
	return s.Save()
}

// mock
type AccountRepository struct {
}

func (a *AccountRepository) Exists(openID string) (bool, error) {
	return false, nil
}

func (a *AccountRepository) Create(openID string) error {
	return nil
}

func (a *AccountRepository) Find(openID string) (int, error) {
	return 1, nil
}
