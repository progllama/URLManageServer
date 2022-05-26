package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"url_manager/app/repositories"
	"url_manager/app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var (
	service = new(services.UserService)
	// repo    = new(repositories.UserRepository)
)

func TestUserApiIndexSuccess(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	repo := repositories.NewUserRepositoryMock()
	var mock repositories.UserRepository = repo

	api := NewUserApi(mock, service)
	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestUserApiIndexSuccessAny(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	repo := repositories.NewUserRepositoryMock()
	repo.Create("test-name", "login-id", "password")
	var mock repositories.UserRepository = repo

	api := NewUserApi(mock, service)
	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	want := `["name":"test-name"}]`
	assert.Equal(t, want, w.Body.String())
}

func TestUserApiIndexFail(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	repo := repositories.NewUserRepositoryMock()
	repo.Error = errors.New("Dummy error.")

	var mock repositories.UserRepository = repo
	api := NewUserApi(mock, service)

	router.GET(route, api.Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", route, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func NewTestRouter() *gin.Engine {
	gin.DefaultWriter = ioutil.Discard
	return gin.New()
}

func TestNewUserApi(t *testing.T) {
	type args struct {
		r repositories.UserRepository
		s services.UserService
	}
	tests := []struct {
		name string
		args args
		want *userApi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserApi(tt.args.r, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userApi_Index(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &userApi{
				repo: tt.fields.repo,
			}
			api.Index(tt.args.ctx)
		})
	}
}

func Test_userApi_Show(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &userApi{
				repo: tt.fields.repo,
			}
			api.Show(tt.args.ctx)
		})
	}
}

func Test_userApi_Create(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &userApi{
				repo: tt.fields.repo,
			}
			api.Create(tt.args.ctx)
		})
	}
}

func Test_userApi_Update(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &userApi{
				repo: tt.fields.repo,
			}
			api.Update(tt.args.ctx)
		})
	}
}

func Test_userApi_Delete(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &userApi{
				repo: tt.fields.repo,
			}
			api.Delete(tt.args.ctx)
		})
	}
}
