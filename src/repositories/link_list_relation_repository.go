package repositories

import (
	"url_manager/models"

	"gorm.io/gorm"
)

type LinkListRelationRepository interface {
	All() ([]models.LinkListRelation, error)
	Find(int) (models.LinkListRelation, error)
	FindByUserId(int) ([]models.LinkListRelation, error)
	FindByParentId(int) ([]models.LinkListRelation, error)
	Add(models.LinkListRelation) error
	Update(models.LinkListRelation) error
	Remove(int) error
	Exists(int) (bool, error)
}

type linkListRelationRepository struct {
	db *gorm.DB
}

func NewLinkListRelationRepository(db *gorm.DB) LinkListRelationRepository {
	return &linkListRelationRepository{
		db,
	}
}

func (repo *linkListRelationRepository) All() ([]models.LinkListRelation, error) {
	var relations []models.LinkListRelation
	err := repo.db.Find(&relations).Error
	return relations, err
}

func (repo *linkListRelationRepository) Find(id int) (models.LinkListRelation, error) {
	var relation models.LinkListRelation
	err := repo.db.Where("id=?", id).First(&relation).Error
	return relation, err
}

func (repo *linkListRelationRepository) FindByUserId(id int) ([]models.LinkListRelation, error) {
	var relations []models.LinkListRelation
	err := repo.db.Where("user_id=?", id).Find(&relations).Error
	return relations, err
}

func (repo *linkListRelationRepository) FindByParentId(id int) ([]models.LinkListRelation, error) {
	var relations []models.LinkListRelation
	err := repo.db.Where("parent_id=?", id).Find(&relations).Error
	return relations, err
}

func (repo *linkListRelationRepository) Add(relation models.LinkListRelation) error {
	result := repo.db.Create(&relation)
	return result.Error
}

func (repo *linkListRelationRepository) Update(relation models.LinkListRelation) error {
	result := repo.db.Save(&relation)
	return result.Error
}

func (repo *linkListRelationRepository) Remove(id int) error {
	result := repo.db.Delete(models.LinkListRelation{}, id)
	return result.Error
}

func (repo *linkListRelationRepository) Exists(id int) (bool, error) {
	var relation models.LinkListRelation
	result := repo.db.Where("id=?", id).Limit(1).Find(&relation)
	return result.RowsAffected > 0, result.Error
}
