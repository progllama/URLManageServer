package sessions

import (
	"reflect"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TestRedisSessionFactory_Create(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		factory *RedisSessionFactory
		args    args
		want    Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &RedisSessionFactory{}
			if got := factory.Create(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisSessionFactory.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisSession_HasUserId(t *testing.T) {
	type fields struct {
		s sessions.Session
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RedisSession{
				s: tt.fields.s,
			}
			if got := rs.HasUserId(); got != tt.want {
				t.Errorf("RedisSession.HasUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisSession_SetUserId(t *testing.T) {
	type fields struct {
		s sessions.Session
	}
	type args struct {
		userId int
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
			rs := &RedisSession{
				s: tt.fields.s,
			}
			rs.SetUserId(tt.args.userId)
		})
	}
}

func TestRedisSession_GetUserId(t *testing.T) {
	type fields struct {
		s sessions.Session
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RedisSession{
				s: tt.fields.s,
			}
			if got := rs.GetUserId(); got != tt.want {
				t.Errorf("RedisSession.GetUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisSession_Clear(t *testing.T) {
	type fields struct {
		s sessions.Session
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RedisSession{
				s: tt.fields.s,
			}
			rs.Clear()
		})
	}
}

func TestRedisSession_getUserIdKeyName(t *testing.T) {
	type fields struct {
		s sessions.Session
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RedisSession{
				s: tt.fields.s,
			}
			if got := rs.getUserIdKeyName(); got != tt.want {
				t.Errorf("RedisSession.getUserIdKeyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRedisSession(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *RedisSession
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisSession(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
