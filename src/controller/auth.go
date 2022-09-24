package controller

import (
	"net/http"
	"url_manager/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/sign_up.tmpl", emptyMap())
}

func SignUp(c *gin.Context) {
	var form SignUpForm
	c.Bind(&form)

	err := repository.CreateUser(form.LoginId, form.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	s := sessions.Default(c)
	s.Set("loginId", form.LoginId)
	err = s.Save()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.tmpl", emptyMap())
}

func Login(c *gin.Context) {
	var form LoginForm
	c.Bind(&form)

	u, err := repository.GetUser(form.LoginId, form.Password)
	if err != nil {
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}

	s := sessions.Default(c)
	s.Set("loginId", u.LoginId)
	err = s.Save()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.Redirect(http.StatusFound, "/")
}
