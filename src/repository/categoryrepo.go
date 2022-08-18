package repository

import "url_manager/model"

type CategoryRepository struct {
}

func (repo *CategoryRepository) All(userID string) []model.Category {
	panic("")
}

func (repo *CategoryRepository) Get(id int) model.Category {
	panic("")
}

func (repo *CategoryRepository) Create(c model.Category) model.Category {
	panic("")
}

func (repo *CategoryRepository) Update(id int, c model.Category) model.Category {
	panic("")
}

func (repo *CategoryRepository) Delete(id int) {
	panic("")
}
