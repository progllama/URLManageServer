package repositories

import (
	"fmt"
	"log"
	"url_manager/app/models"
	"url_manager/db"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (repo *UserRepository) All() ([]models.User, error) {
	var users []models.User
	dbResult := repo.db.Table("users").Select("*").Find(&users)

	if dbResult.Error == nil {
		return users, nil
	} else {
		return []models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) Find(condition models.User) ([]models.User, error) {
	var users []models.User
	dbResult := repo.db.Where(&condition).Find(&users)

	if dbResult.Error == nil {
		return users, nil
	} else {
		return []models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) FindById(id int) (models.User, error) {
	fmt.Println(id)
	var user models.User
	dbResult := db.GetDB().Where("id=?", id).First(&user)

	log.Println(id)

	if dbResult.Error == nil {
		return user, nil
	} else {
		return models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	dbResult := db.GetDB().Where("name=?", name).First(&user)

	if dbResult.Error == nil {
		return user, nil
	} else {
		return models.User{}, dbResult.Error
	}
}

func (repo *UserRepository) FindByLoginId(loginId string) (models.User, error) {
	var user models.User
	dbResult := db.GetDB().Where("login_id=?", loginId).First(&user)

	log.Println(db.GetDB().Model(&models.User{}))

	if dbResult.Error == nil {
		return user, nil
	} else {
		return models.User{}, dbResult.Error
	}
}

type SafeUser struct {
	Name string
	ID   string
}

func (repo *UserRepository) AllIdAndNames() ([]SafeUser, error) {
	var users []SafeUser
	dbResult := repo.db.Table("users").Find(&users)

	if dbResult.Error != nil {
		return make([]SafeUser, 0), dbResult.Error
	}

	return users, nil
}

func (repo *UserRepository) Create(name string, loginId string, password string) error {
	user := models.User{
		Name:    name,
		LoginId: loginId,
	}
	hashPass, err := user.GenerateHashFromPassword(password)
	if err != nil {
		log.Println("Fail to generate hash.")
		return err
	}
	user.Password = hashPass

	dbResult := repo.db.Create(&user)
	return dbResult.Error
}

func (repo *UserRepository) Update(id int, name string, loginId string, password string) error {
	u := models.User{
		ID:       id,
		Name:     name,
		LoginId:  loginId,
		Password: password,
	}
	dbResult := repo.db.Updates(u)

	if dbResult != nil {
		return dbResult.Error
	} else {
		return nil
	}
}

func (repo *UserRepository) Delete(id int) error {
	var user models.User
	dbResult := repo.db.Where("id=?", id).Delete(&user)
	return dbResult.Error
}

func (repo *UserRepository) Exists(condition models.User) (bool, error) {
	var users []models.User
	log.Println(condition)
	dbResult := repo.db.Where(&condition).Limit(1).First(&users)
	if dbResult.Error != nil {
		return false, dbResult.Error
	}

	if len(users) > 0 {
		log.Println("records > 0")
		return true, nil
	} else {
		log.Println("record <= 0")
		log.Println(len(users))
		return false, nil
	}
}
