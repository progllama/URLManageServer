package auth

// Google, Twitter, Facebook
// などの外部アカウントを表現するためのクラス。
type User interface {
	LoginId() string
}

type UserService interface {
	Fetch() (User, error)
}

type UserServiceFactory interface {
	NewUserService() UserService
}
