package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var TwitterConf = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login",
	Scopes:      []string{"users.read", "tweet.read"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://twitter.com/i/oauth2/authorize",
		TokenURL: "https://api.twitter.com/2/oauth2/token",
	},
}

func TwitterAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)

		ss := s.Get("twitter state")
		qs := c.Query("state")

		if ss == nil || ss != qs {
			return
		}

		code := c.Query("code")
		verifier := s.Get("verifier")
		if verifier == nil {
			return
		}
		token, err := TwitterConf.Exchange(
			context.TODO(),
			code,
			oauth2.SetAuthURLParam("code_verifier", verifier.(string)),
		)
		if err != nil {
			return
		}

		client := oauth2.NewClient(
			context.TODO(),
			oauth2.StaticTokenSource(token),
		)

		res, err := client.Get("https://api.twitter.com/2/users/me")
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
func TwitterLoginURL(state string, codeChallenge string) string {
	return GoogleConf.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
}
