package api

import (
	"github.com/gin-gonic/gin"
	goauth "google.golang.org/api/oauth2/v2"
)

func getUserInfo(ctx *gin.Context) goauth.Userinfo {
	var (
		res goauth.Userinfo
		ok  bool
	)

	val := ctx.MustGet("user")
	if res, ok = val.(goauth.Userinfo); !ok {
		res = goauth.Userinfo{Name: "no user"}
	}

	return res
}
