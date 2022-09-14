package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var FacebookConf = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login",
	Scopes:      []string{"public_profile"},
	Endpoint:    facebook.Endpoint,
}

func FacebookAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)

		ss := s.Get("facebook state")
		qs := c.Query("state")

		if ss == nil || ss != qs {
			return
		}

		code := c.Query("code")
		token, err := FacebookConf.Exchange(
			context.TODO(),
			code,
		)
		if err != nil {
			return
		}

		client := oauth2.NewClient(
			context.TODO(),
			oauth2.StaticTokenSource(token),
		)

		res, err := client.Get("https://graph.facebook.com/v14.0/me")
		if err != nil {
			return
		}
		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}

		var u struct {
			ID string `json:"id"`
		}
		err = json.Unmarshal(b, &u)
		if err != nil {
			return
		}

		s.Set("id", u.ID)
		s.Save()
	}
}

func FacebookLoginURL(state string) string {
	return GoogleConf.AuthCodeURL(state)
}
