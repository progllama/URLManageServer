package repositories

import (
	"url_manager/domain/models"

	"gorm.io/gorm"
)

type LinkRepository interface {
	All() ([]models.Link, error)
	Find(int) (models.Link, error)
	Add(models.Link) error
	Update(models.Link) error
	Remove(int) error
	Exists(int) (bool, error)
}

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepository{
		db,
	}
}

func (repo *linkRepository) All() ([]models.Link, error) {
	var urls []models.Link
	err := repo.db.Find(&urls).Error
	return urls, err
}

func (repo *linkRepository) Find(id int) (models.Link, error) {
	var url models.Link
	err := repo.db.Where("id=?", id).First(&url).Error
	return url, err
}

func (repo *linkRepository) Add(url models.Link) error {
	result := repo.db.Create(&url)
	return result.Error
}

func (repo *linkRepository) Update(url models.Link) error {
	result := repo.db.Save(&url)
	return result.Error
}

func (repo *linkRepository) Remove(id int) error {
	result := repo.db.Delete(models.Link{}, id)
	return result.Error
}

func (repo *linkRepository) Exists(id int) (bool, error) {
	var url models.Link
	result := repo.db.Where("id=?", id).Limit(1).Find(url)
	return result.RowsAffected > 0, result.Error
}
