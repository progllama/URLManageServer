package services

import "url_manager/domain/models"

// アプリケーションレイヤのサービス
// ドメインの知識を入れてはいけない。
// ドメインレベルの手続きはDomainのサービスに入れる。

type UserService interface {
	FindUsers() UserServiceResponse
	FindUser(string) UserServiceResponse
	// Create(UserCreateRequest) UserServiceResponse
	Create(models.User) UserServiceResponse
	Update(models.User) UserServiceResponse
	Delete(string) UserServiceResponse
}

// FindUserなどもリクエストかすべき？
type UserCreateRequest struct {
	LoginID  string `json:"login_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
}

type UserServiceResponse struct {
	Code int
	Body interface{}
}
