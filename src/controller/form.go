package controller

type SignUpForm struct {
	LoginId  string `form:"loginId"`
	Password string `form:"password"`
}

type LoginForm struct {
	LoginId  string `form:"loginId"`
	Password string `form:"password"`
}

type LinkForm struct {
	Title       string `form:"title"`
	URL         string `form:"url"`
	Description string `form:"description"`
}

type FolderForm struct {
	Title string `form:"title"`
}
