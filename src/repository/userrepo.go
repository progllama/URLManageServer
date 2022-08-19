package repository

import "url_manager/model"

type UserRepository struct {
}

// used login
func (repo *UserRepository) GetByUserId(uid string) model.User {
	db := getDB()
	var user model.User
	db.Where("user_id=?", uid).Find(&user)
	return user
}

func (repo *UserRepository) Create(u model.User) model.User {
	db := getDB()
	db.Create(&u)
	var created model.User
	db.Where("user_id=?", u.UserID).Find(&created)
	return created
}

func (repo *UserRepository) Update(id int, u model.User) model.User {
	db := getDB()
	db.Where("id=?", id).Save(&u)
	var updated model.User
	db.Where("id=?", id).Find(&updated)
	return updated
}

func (repo *UserRepository) Delete(id int) {
	db := getDB()
	u := model.User{}
	u.ID = uint(id)
	db.Delete(&u)
}
