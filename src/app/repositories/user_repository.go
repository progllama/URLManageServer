package repositories

import (
	"errors"
	"url_manager/domain/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	All() ([]models.User, error)
	AllIdName() ([]models.User, error)
	FindById(int) (models.User, error)
	FindByName(string) (models.User, error)
	FindByLoginId(string) (models.User, error)
	Create(name, loginId, password string) error
	Update(id int, name, login, password string) error
	Delete(id int) error
	// HasUserId(int) (bool, error)
	HasUserName(string) (bool, error)
	// HasUserLoginId(string) (bool, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepositoryImplGorm(db *gorm.DB) UserRepository {
	return &gormUserRepository{
		db,
	}
}

func (repo *gormUserRepository) All() ([]models.User, error) {
	var users []models.User
	result := repo.db.Table("users").Select("id, name").Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}

func (repo *gormUserRepository) AllIdName() ([]models.User, error) {
	var users []models.User
	dbResult := repo.db.Table("users").Find(&users)

	if dbResult.Error != nil {
		return make([]models.User, 0), dbResult.Error
	}

	return users, nil
}

func (repo *gormUserRepository) FindById(id int) (models.User, error) {
	var user models.User
	dbResult := repo.db.Where("id=?", id).First(&user)
	if dbResult.Error != nil {
		return models.User{}, dbResult.Error
	}

	return user, nil
}

func (repo *gormUserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	dbResult := repo.db.Where("name=?", name).First(&user)
	if dbResult.Error == nil {
		return models.User{}, dbResult.Error
	}

	return user, nil
}

func (repo *gormUserRepository) FindByLoginId(loginId string) (models.User, error) {
	var user models.User
	dbResult := repo.db.Where("login_id=?", loginId).First(&user)
	if dbResult.Error == nil {
		return models.User{}, dbResult.Error
	}

	return user, nil
}

func (repo *gormUserRepository) Create(name string, loginId string, password string) error {
	user := models.User{
		Name:    name,
		LoginId: loginId,
	}
	hashPass, err := user.GenerateHashFromPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashPass

	dbResult := repo.db.Create(&user)
	return dbResult.Error
}

func (repo *gormUserRepository) Update(id int, name string, loginId string, password string) error {
	u := models.User{
		ID:       id,
		Name:     name,
		LoginId:  loginId,
		Password: password,
	}
	dbResult := repo.db.Model(&models.User{}).Updates(u)

	if dbResult != nil {
		return dbResult.Error
	} else {
		return nil
	}
}

func (repo *gormUserRepository) Delete(id int) error {
	var user models.User
	dbResult := repo.db.Where("id=?", id).Delete(&user)
	return dbResult.Error
}

func (repo *gormUserRepository) HasUserName(name string) (bool, error) {
	var users []models.User
	dbResult := repo.db.Select("id").Where("name=?", name).Limit(1).Take(&users)
	if !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) && dbResult.Error != nil {
		return false, dbResult.Error
	}

	if len(users) == 0 {
		return false, nil
	}

	return true, nil
}
