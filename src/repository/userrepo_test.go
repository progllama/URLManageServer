package repository_test

import (
	"errors"
	"testing"
	"url_manager/model"
	"url_manager/repository"
)

func TestUserRepoAll(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatal(err)
	}
	con, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal(err)
	}

	fixture := []model.User{
		{Name: "Alice"},
		{Name: "Jinx"},
		{Name: "John"},
		{Name: "Finn"},
	}
	if err = db.CreateInBatches(fixture, 100).Error; err != nil {
		t.Fatal(err)
	}
	remove := &model.User{}
	remove.ID = 4
	db.Delete(remove)

	repo := repository.NewUserRepository(db)
	users, err := repo.All()

	if err != nil {
		t.Error(err)
	}

	if len(users) != len(fixture)-1 {
		t.Error("num of db item is different")
	}

	for i, u := range users {
		if fixture[i].Name != u.Name {
			t.Errorf("catch non match item, %s, %s", fixture[i].Name, u.Name)
		}
		if !u.CreatedAt.IsZero() {
			t.Error("get non selected column item")
		}
		if !u.UpdatedAt.IsZero() {
			t.Error("get non selected column item")
		}
	}
}

func TestNormalCaseUserRepoGet(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatal(err)
	}
	con, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal(err)
	}

	fixture := []model.User{
		{Name: "Alice"},
		{Name: "John"},
		{Name: "Finn"},
	}
	if err = db.CreateInBatches(fixture, 100).Error; err != nil {
		t.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	user, err := repo.Get(2)

	if err != nil {
		t.Error(err)
	}
	if user.ID != 2 {
		t.Error("catch non match item")
	}
	if fixture[1].Name != user.Name {
		t.Error("catch non match item")
	}
	if !user.CreatedAt.IsZero() {
		t.Error("get non selected column item")
	}
	if !user.UpdatedAt.IsZero() {
		t.Error("get non selected column item")
	}
}

func TestErrorCaseUserRepoGet(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatal(err)
	}
	con, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal(err)
	}

	fixture := []model.User{
		{Name: "Alice"},
		{Name: "John"},
		{Name: "Finn"},
	}
	if err = db.CreateInBatches(fixture, 100).Error; err != nil {
		t.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	user, err := repo.Get(4)

	if err == nil {
		t.Error("should error")
	}
	if !errors.Is(err, repository.ErrItemNotFound) {
		t.Error("unexpected error returned")
	}
	if user.ID != 0 {
		t.Error("catch non match item")
	}
}

func TestNormalCaseUserRepoCreate(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatal(err)
	}
	con, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal(err)
	}

	user := model.User{
		Name: "Alice$%",
	}
	repo := repository.NewUserRepository(db)
	created, err := repo.Create(user)
	if err != nil {
		t.Error(err)
	}
	if created.Name != user.Name {
		t.Error("different name")
	}
	if !created.CreatedAt.IsZero() {
		t.Error("select fail")
	}
}

func TestNormalCaseUserRepoUpdate(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatal(err)
	}
	con, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal(err)
	}

	user := model.User{
		Name: "Alice$%",
	}
	repo := repository.NewUserRepository(db)
	repo.Create(user)
	if err != nil {
		t.Error(err)
	}

	user.ID = 1
	user.Name = "Alice"
	updated, err := repo.Update(user)
	if err != nil {
		t.Error(err)
	}
	if updated.Name != "Alice" {
		t.Error(err)
	}

}

func TestNormalCaseUserRepoDelete(t *testing.T) {
	db, err := createTestDB()
	if err != nil {
		t.Fatal(err)
	}
	con, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal(err)
	}

	user := model.User{
		Name: "Alice$%",
	}
	repo := repository.NewUserRepository(db)
	repo.Create(user)
	if err != nil {
		t.Error(err)
	}

	err = repo.Delete(1)
	if err != nil {
		t.Error(err)
	}
	users, err := repo.All()
	if err != nil {
		t.Error(err)
	}
	if len(users) != 0 {
		t.Error("fail to delete")
	}
}
