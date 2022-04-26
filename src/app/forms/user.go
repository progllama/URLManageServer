package forms

type UserCreateForm struct {
	Name     string `form:"name"`
	LoginId  string `form:"login_id"`
	Password string `form:"password"`
}

type UserEditForm struct {
	Name     string `form:"name"`
	LoginId  string `form:"login_id"`
	Password string `form:"password"`
}
