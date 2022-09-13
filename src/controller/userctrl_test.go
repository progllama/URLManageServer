package controllers

import (
	"reflect"
	"testing"
	"url_manager/repository"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.ctx)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUser(tt.args.ctx)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUser(tt.args.ctx)
		})
	}
}

func Test_getUserRepo(t *testing.T) {
	tests := []struct {
		name string
		want *repository.UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
