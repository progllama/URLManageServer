package api

import (
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
	service services.UserService
	config  UserApiConfig
)

func TestUserApiIndexSuccess(t *testing.T) {
	router := NewTestRouter()
	route := "/api/users"

	api := NewUserApi(service, config)
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

	api := NewUserApi(service, config)
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

	api := NewUserApi(service, config)

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
		s services.UserService
		c UserApiConfig
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
			if got := NewUserApi(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
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
			api := &userApi{}
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
			api := &userApi{}
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
			api := &userApi{}
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
			api := &userApi{}
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
			api := &userApi{}
			api.Delete(tt.args.ctx)
		})
	}
}
