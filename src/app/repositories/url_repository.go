package repositories

import (
	"url_manager/app/models"
	"url_manager/db"
)

type URLRepository interface {
	GetByUserID(int) ([]models.URL, error)
	GetByID(int) (models.URL, error)
	Create(models.URL) error
	Update(models.URL) error
	Destory(models.URL) error
}

type PostgreSQLURLRepository struct {
}

func (repo PostgreSQLURLRepository) GetAll() ([]models.URL, error) {
	db := db.GetDB()
	if db.Error != nil {
		return []models.URL{}, db.Error
	}

	var urls []models.URL
	if err := db.Select("*").Find(&urls).Error; err != nil {
		return []models.URL{}, err
	}

	return urls, nil
}

func (repo PostgreSQLURLRepository) GetByUserID(id int) ([]models.URL, error) {
	db := db.GetDB()
	if db.Error != nil {
		return []models.URL{}, db.Error
	}

	var urls []models.URL
	if err := db.Select("id, url, title").Where("user_id=?", id).Find(&urls).Error; err != nil {
		return []models.URL{}, err
	}

	return urls, nil
}

func (repo PostgreSQLURLRepository) GetByID(id int) (models.URL, error) {
	db := db.GetDB()
	if db.Error != nil {
		return models.URL{}, db.Error
	}

	var url models.URL
	err := db.Select("title, url").Where("id=?", id).Find(&url).Error
	if err != nil {
		return models.URL{}, err
	}

	return url, nil
}

func (repo PostgreSQLURLRepository) Create(url models.URL) error {
	db := db.GetDB()
	if db.Error != nil {
		return db.Error
	}

	g := db.Create(&url)
	if g.Error != nil {
		return g.Error
	}

	return nil
}

func (repo PostgreSQLURLRepository) Update(url models.URL) error {
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

func (repo PostgreSQLURLRepository) Delete(url models.URL) error {
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
