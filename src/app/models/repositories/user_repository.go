package repositories

import (
	"url_manager/app/models"
	"url_manager/db"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Userをそのまま返しているのでテーブル構造がバレてしまう。

func GetAllUsers() ([]models.User, error) {
	// Retrieve all users.
	var users []models.User
	// TODO 命名: reslutだと取得結果だと思ってそれを返すかも。
	result := getDB().Table("users").Select("name, id").Find(&users)

	// Handle err.
	// NOTE ifで処理ではなく値が切り替わるときにdryじゃない気がする。
	if result.Error == nil {
		return users, result.Error
	} else {
		return users, result.Error
	}
}

func GetUserByID(id int) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	db.Table("users").Where("id = ?", id).First(&user)

	return user, nil
}

func GetUserByName(name string) (models.User, error) {
	db := db.GetDB()
	var user models.User

	if err := db.Table("users").Where("name=?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) error {
	db := db.GetDB()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12) // 2 ^ 12 回　ストレッチ回数

	if err != nil {
		return err
	}

	user.Password = string(hashedPass)

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserByID(id int, c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	user.ID = uint(id)
	db.Save(&user)

	return user, nil
}

func DeleteUserByID(id int) error {
	db := db.GetDB()
	var user models.User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func HasUser(name string) (bool, error) {
	db := db.GetDB()
	var users []models.User
	err := db.Select("count(*) > 0").Where("name=?", name).Limit(1).Find(&users).Error
	return len(users) > 0, err
}

func getDB() *gorm.DB {
	return db.GetDB()
}
