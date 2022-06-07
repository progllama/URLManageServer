package models

type Credential struct {
	Id       string `json:"email"`
	Password string `json:"password"`
}
