package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginIdProviderFactory interface {
	NewProvider(*gin.Context) LoginIdProvider
}

type LoginIdProvider interface {
	LoginId() string
}

func BlockAnonymousUsers(factory LoginIdProviderFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		provider := factory.NewProvider(c)

		account := NewAccount()
		account.SetLoginID(provider.LoginId())
		if !account.Exists() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func BlockUsersExceptOwner(factory LoginIdProviderFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		service := NewLoginService(c)
		provider := factory.NewProvider(c)

		if !service.IsLogin(provider.LoginId()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
