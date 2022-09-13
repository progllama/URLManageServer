package auth

import "golang.org/x/oauth2"

// Google, Twitter, Facebook
// などの外部アカウントを表現するためのクラス。
type User interface {
	LoginId() string
}

type UserService interface {
	Fetch(token *oauth2.Token) (User, error)
}

type UserServiceFactory interface {
	NewUserService() UserService
}
