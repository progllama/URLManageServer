package repositories

import (
	"url_manager/app/models"

	"gorm.io/gorm"
)

type UrlRepository interface {
	All() ([]models.Url, error)
	Find(int) (models.Url, error)
	Add(models.Url) error
	Update(models.Url) error
	Remove(int) error
	Exists(int) (bool, error)
}

type urlRepositoryImplPostgres struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepositoryImplPostgres{
		db,
	}
}

func (repo *urlRepositoryImplPostgres) All() ([]models.Url, error) {
	var urls []models.Url
	err := repo.db.Find(&urls).Error
	return urls, err
}

func (repo *urlRepositoryImplPostgres) Find(id int) (models.Url, error) {
	var url models.Url
	err := repo.db.Where("id=?", id).First(&url).Error
	return url, err
}

func (repo *urlRepositoryImplPostgres) Add(url models.Url) error {
	result := repo.db.Create(&url)
	return result.Error
}

func (repo *urlRepositoryImplPostgres) Update(url models.Url) error {
	result := repo.db.Save(&url)
	return result.Error
}

func (repo *urlRepositoryImplPostgres) Remove(id int) error {
	result := repo.db.Delete(models.Url{}, id)
	return result.Error
}

func (repo *urlRepositoryImplPostgres) Exists(id int) (bool, error) {
	var url models.Url
	result := repo.db.Where("id=?", id).Limit(1).Find(url)
	return result.RowsAffected > 0, result.Error
}
