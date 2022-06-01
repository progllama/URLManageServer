package repositories

import (
	"fmt"
	"log"
	"url_manager/db"
	"url_manager/domain/models"
)

// type URLRepository interface {
// 	FindByUserID(int) ([]models.URL, error)
// 	FindByID(int) (models.URL, error)
// 	Create(models.URL) error
// 	Update(models.URL) error
// 	Destory(models.URL) error
// }

func NewUrlRepository() *PostgreSQLURLRepository {
	return &PostgreSQLURLRepository{}
}

type PostgreSQLURLRepository struct {
}

func (repo PostgreSQLURLRepository) All() ([]models.Url, error) {
	db := db.GetDB()
	if db.Error != nil {
		return []models.Url{}, db.Error
	}

	var urls []models.Url
	if err := db.Select("*").Find(&urls).Error; err != nil {
		return []models.Url{}, err
	}

	return urls, nil
}

func (repo PostgreSQLURLRepository) FindByUserID(id int) ([]models.Url, error) {
	db := db.GetDB()
	if db.Error != nil {
		return []models.Url{}, db.Error
	}

	var user models.User
	user.ID = id
	var urls []models.Url
	if err := db.Model(&user).Association("Urls").Find(&urls); err != nil {
		fmt.Println(err)
		return []models.Url{}, err
	}

	return urls, nil
}

func (repo PostgreSQLURLRepository) FindByID(id int) (models.Url, error) {
	db := db.GetDB()
	if db.Error != nil {
		return models.Url{}, db.Error
	}

	var url models.Url
	err := db.Select("title, url").Where("id=?", id).Find(&url).Error
	if err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func (repo PostgreSQLURLRepository) Create(ownerId int, url *models.Url) error {
	db := db.GetDB()
	if db.Error != nil {
		return db.Error
	}

	user := models.User{}
	user.ID = ownerId
	err := db.Model(&user).Association("Urls").Append(&[]models.Url{*url})
	if err != nil {
		log.Println("ERR ", err)
		return err
	}

	return nil
}

func (repo PostgreSQLURLRepository) Update(url models.Url) error {
	db := db.GetDB()
	if db.Error != nil {
		return db.Error
	}

	g := db.Updates(&url)
	if g.Error != nil {
		return g.Error
	}

	return nil
}

func (repo PostgreSQLURLRepository) Delete(url models.Url) error {
	db := db.GetDB()
	if db.Error != nil {
		return db.Error
	}

	g := db.Delete(&url)
	if g.Error != nil {
		return g.Error
	}

	return nil
}
