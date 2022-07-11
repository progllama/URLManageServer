package api

import (
	"errors"
	"url_manager/models"
	"url_manager/repositories"

	"github.com/gin-gonic/gin"
	goauth "google.golang.org/api/oauth2/v2"
	"gorm.io/gorm"
)

func SignIn(repo repositories.UserRepository) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var res *goauth.Userinfo

		if v, ok := ctx.Get("user"); ok {
			res = v.(*goauth.Userinfo)
		} else {
			res = &goauth.Userinfo{Name: "no user"}
		}

		_, err := repo.FindByOpenID(res.Id)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			repo.Add(models.User{OpenID: res.Id})
		}
		ctx.Redirect(302, "/home")
	}
}
