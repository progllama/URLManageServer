package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url_manager/database"
	"url_manager/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestGetLinkList(t *testing.T) {
	gin.DefaultWriter = io.Discard
	router := gin.New()
	router.Use(sessions.Sessions("test", cookie.NewStore([]byte("test-secret"))))

	api := NewLinksApi(
		repositories.NewUserRepository(database.Database()),
		repositories.NewLinkRepository(database.Database()),
		repositories.NewLinkListRepository(database.Database()),
		repositories.NewLinkListRelationRepository(database.Database()),
	)
	router.GET("/users/:user_id/links", api.Index)

	w := httptest.NewRecorder()
	body := ""
	req, _ := http.NewRequest(http.MethodGet, "/users/1/links/1", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
}
