package repository

import "url_manager/model"

type LinkRepository struct {
}

func (repo *LinkRepository) All(userID int) []model.Link {
	db := getDB()
	u := model.NewUser(userID)
	var links []model.Link
	db.Model(u).Association("Links").Find(&links)
	return links
}

func (repo *LinkRepository) Create(userID int, l model.Link) model.Link {
	db := getDB()
	u := model.NewUser(userID)
	db.Model(u).Association("Links").Append(&l)
	return l
}

func (repo *LinkRepository) Update(id int, l model.Link) model.Link {
	db := getDB()
	db.Where("id", id).Save(l)
	return l
}

func (repo *LinkRepository) Delete(id int) {
	db := getDB()
	l := model.NewLink(id)
	db.Delete(l)
}
