package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	params := HomeParams{}
	params.IsLogin = c.GetBool("isLogin")
	params.LoginId = c.GetString("loginId")

	c.HTML(http.StatusOK, "static/home.tmpl", params)
}

type HomeParams struct {
	CommonParams
	LoginId string
}
