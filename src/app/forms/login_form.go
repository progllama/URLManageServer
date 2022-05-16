package forms

type LoginForm struct {
	LoginId  string `form:"login_id"`
	Password string `form:"password"`
}
