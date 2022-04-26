package repositories

import (
	"fmt"
	"log"
	"url_manager/app/models"
	"url_manager/db"
)

// type URLRepository interface {
// 	GetByUserID(int) ([]models.URL, error)
// 	GetByID(int) (models.URL, error)
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
	fmt.Println(id)
	fmt.Println(db.Model(&user).Error)
	fmt.Println(db.Model(&user).Where("id=?", id).Error)
	fmt.Println(db.Model(&user).Where("id=?", id).Association("Urls").Find(&urls).Error)
	if err := db.Model(&user).Association("Urls").Find(&urls).Error; err != nil {
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
	g := db.Model(&user).Association("Urls").Append(&[]models.Url{*url})
	if g.Error != nil {
		log.Println("ERR ", g.Error)
		return g.Error
	}

	return nil
}

func (repo PostgreSQLURLRepository) Update(url models.Url) error {
	db := db.GetDB()
	if db.Error != nil {
		return db.Error
	}

	g := db.Update(&url)
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
