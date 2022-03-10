package repositories

import (
	"url_manager/app/models"
	"url_manager/db"
)

type URLRepository interface {
	GetByUserID(int) ([]models.URL, error)
	Create(models.URL) error
	Update(models.URL) error
	Destory(models.URL) error
}

type DefaultURLRepositoryImpl struct {
}

func (self DefaultURLRepositoryImpl) GetByUserID(id int) ([]models.URL, error) {
	db := db.GetDB()
	if db.Error != nil {
		return []models.URL{}, db.Error
	}

	var urls []models.URL
	if err := db.Select("url, title").Where("UserID=?", id).Find(&urls).Error; err != nil {
		return []models.URL{}, err
	}

	return urls, nil
}

func (self DefaultURLRepositoryImpl) Create(url models.URL) error {
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

func (self DefaultURLRepositoryImpl) Update(url models.URL) error {
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

func (self DefaultURLRepositoryImpl) Destroy(url models.URL) error {
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
