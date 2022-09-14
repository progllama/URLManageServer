package auth

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var GoogleConf = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login",
	Scopes:      []string{"public_profile"},
	Endpoint:    google.Endpoint,
}

func GoogleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)

		ss := s.Get("google state")
		qs := c.Query("state")

		if ss == nil || ss != qs {
			return
		}

		code := c.Query("code")
		token, err := GoogleConf.Exchange(
			context.TODO(),
			code,
		)
		if err != nil {
			return
		}

		service, err := goauth.NewService(
			context.TODO(),
			option.WithTokenSource(GoogleConf.TokenSource(context.TODO(), token)))
		if err != nil {
			return
		}

		userInfo, err := service.Userinfo.Get().Do()
		if err != nil {
			return
		}

		s.Set("id", userInfo.Id)
		s.Save()
	}
}

func GoogleLoginURL(state string) string {
	return GoogleConf.AuthCodeURL(state)
}
