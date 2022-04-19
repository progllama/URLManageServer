package forms

type UrlCreateForm struct {
	Url         string `form:"url"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Note        string `form:"note"`
}

type UrlEditForm struct {
}
