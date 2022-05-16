package middlewares

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewLoginRequireMiddleware(t *testing.T) {
	// gin.DefaultWriter = io.Discard
	// router := gin.New()

	type args struct {
		factory   SessionFactory
		onSuccess func(*gin.Context)
		onFail    func(*gin.Context)
	}
	tests := []struct {
		name string
		args args
		want *LoginRequireMiddleWare
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginRequireMiddleware(tt.args.factory, tt.args.onSuccess, tt.args.onFail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginRequireMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginRequireMiddleWare_Handler(t *testing.T) {
	type fields struct {
		factory   SessionFactory
		onSuccess func(*gin.Context)
		onFail    func(*gin.Context)
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := &LoginRequireMiddleWare{
				factory:   tt.fields.factory,
				onSuccess: tt.fields.onSuccess,
				onFail:    tt.fields.onFail,
			}
			if got := mw.Handler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginRequireMiddleWare.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginRequireMiddleWare_getSession(t *testing.T) {
	type fields struct {
		factory   SessionFactory
		onSuccess func(*gin.Context)
		onFail    func(*gin.Context)
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := &LoginRequireMiddleWare{
				factory:   tt.fields.factory,
				onSuccess: tt.fields.onSuccess,
				onFail:    tt.fields.onFail,
			}
			if got := mw.getSession(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginRequireMiddleWare.getSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
