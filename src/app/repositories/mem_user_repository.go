package repositories

import "url_manager/app/models"

type MemUserRepository struct {
	users []models.User
}

func (repo MemUserRepository) All() (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) AllIdName() (SafeUser, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) FindById(_ int) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) FindByName(_ string) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) FindByLoginId(_ string) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) Create(name string, loginId string, password string) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) Update(id int, name string, login string, password string) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) Delete(id int) error {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) HasUserId(_ int) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) HasUserName(_ string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (repo MemUserRepository) HasUserLoginId(_ string) (bool, error) {
	panic("not implemented") // TODO: Implement
}
