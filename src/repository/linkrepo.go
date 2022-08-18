package repository

import "url_manager/model"

type LinkRepository struct {
}

func (repo *LinkRepository) All(userID string) []model.Link {
	panic("")
}

func (repo *LinkRepository) Get(id int) model.Link {
	panic("")
}

func (repo *LinkRepository) Create(c model.Link) model.Link {
	panic("")
}

func (repo *LinkRepository) Update(id int, c model.Link) model.Link {
	panic("")
}

func (repo *LinkRepository) Delete(id int) {
	panic("")
}
