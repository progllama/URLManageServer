package repositories

import (
	"url_manager/models"

	"gorm.io/gorm"
)

type LinkListRepository interface {
	All() ([]models.LinkList, error)
	Find(int) (models.LinkList, error)
	FindByUserId(int) ([]models.LinkList, error)
	Add(models.LinkList) error
	Update(models.LinkList) error
	Remove(int) error
	Exists(int) (bool, error)
}

type linkListRepository struct {
	db *gorm.DB
}

func NewLinkListRepository(db *gorm.DB) LinkListRepository {
	return &linkListRepository{
		db,
	}
}

func (repo *linkListRepository) All() ([]models.LinkList, error) {
	var lists []models.LinkList
	err := repo.db.Find(&lists).Error
	return lists, err
}

func (repo *linkListRepository) Find(id int) (models.LinkList, error) {
	var list models.LinkList
	err := repo.db.Where("id=?", id).First(&list).Error
	return list, err
}

func (repo *linkListRepository) FindByUserId(id int) ([]models.LinkList, error) {
	var lists []models.LinkList
	err := repo.db.Where("user_id=?", id).Find(&lists).Error
	return lists, err
}

func (repo *linkListRepository) Add(list models.LinkList) error {
	result := repo.db.Create(&list)
	return result.Error
}

func (repo *linkListRepository) Update(list models.LinkList) error {
	result := repo.db.Save(&list)
	return result.Error
}

func (repo *linkListRepository) Remove(id int) error {
	result := repo.db.Delete(&models.LinkList{}, id)
	return result.Error
}

func (repo *linkListRepository) Exists(id int) (bool, error) {
	var list models.LinkList
	result := repo.db.Where("id=?", id).Limit(1).Find(list)
	return result.RowsAffected > 0, result.Error
}
