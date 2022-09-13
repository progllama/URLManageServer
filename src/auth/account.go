package auth

import (
	"time"

	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Account interface {
	GetID() int
	SetID(int)
	GetLoginID() string
	SetLoginID(string)
	GetAccessToken() string
	SetAccessToken(string)
	GetRefreshToken() string
	SetRefreshToken(string)
	GetTokenExpire() time.Time
	SetTokenExpire(time.Time)
	Create() error
	Exists() bool
}

// TODO: 使用するアカウント構造体の生成もFactoryで抽象化する
func NewAccount() Account {
	return &AccountModel{}
}

type AccountModel struct {
	gorm.Model
	LoginID      string `gorm:"unique"`
	AccessToken  string
	RefreshToken string
	Expire       time.Time
}

func (account *AccountModel) GetID() int {
	return int(account.ID)
}

func (account *AccountModel) SetID(id int) {
	account.ID = uint(id)
}

func (account *AccountModel) GetLoginID() string {
	return account.LoginID
}

func (account *AccountModel) SetLoginID(loginId string) {
	account.LoginID = loginId
}

func (account *AccountModel) GetAccessToken() string {
	return account.AccessToken
}

func (account *AccountModel) SetAccessToken(accessToken string) {
	account.AccessToken = accessToken
}

func (account *AccountModel) GetRefreshToken() string {
	return account.RefreshToken
}

func (account *AccountModel) SetRefreshToken(refreshToken string) {
	account.AccessToken = refreshToken
}

func (account *AccountModel) GetTokenExpire() time.Time {
	return account.Expire
}

func (account *AccountModel) SetTokenExpire(expire time.Time) {
	account.Expire = expire
}

func (account *AccountModel) Exists() bool {
	result := DB.Where("login_id=?", account.LoginID).First(account)
	return result.Error == nil
}

func (account *AccountModel) Create() error {
	result := DB.Create(account)
	return result.Error
}

func AssociateAccount(user User, token *oauth2.Token) error {
	account := NewAccount()
	account.SetLoginID(user.LoginId())
	if account.Exists() {
		return nil
	}

	account.SetAccessToken(token.AccessToken)
	account.SetRefreshToken(token.RefreshToken)
	account.SetTokenExpire(token.Expiry)
	return account.Create()
}
