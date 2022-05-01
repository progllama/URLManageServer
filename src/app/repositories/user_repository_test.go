package repositories

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"url_manager/app/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// TODO
// use txdb sql-mock gomock

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

type Any struct{}

func (a Any) Match(v driver.Value) bool {
	return true
}

func TestAll(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "login_id"}).AddRow("2222", "yuta", "yuta-id"))

	repo := UserRepository{
		db: db,
	}
	users, err := repo.All()
	t.Log(users)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Fatal(err)
	}
}

func TestFindByName(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "login_id"}).
			AddRow("1", "user1", "i_am_user1").
			AddRow("2", "user2", "i_am_user2"))

	repo := UserRepository{
		db: db,
	}

	user, err := repo.FindByName("user2")
	if err != nil {
		t.Error(err)
	}

	if user.Name != "user2" {
		t.Error(err)
	}
}

func TestFindById(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "login_id"}).
			AddRow("1", "user1", "i_am_user1").
			AddRow("2", "user2", "i_am_user2"))

	repo := UserRepository{
		db: db,
	}

	user, err := repo.FindById(2)
	if err != nil {
		t.Error(err)
	}

	if user.Name != "user2" {
		t.Error(err)
	}
}

func TestFindByLoginId(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "login_id"}).
			AddRow("1", "user1", "i_am_user1").
			AddRow("2", "user2", "i_am_user2"))

	repo := UserRepository{
		db: db,
	}

	user, err := repo.FindByLoginId("user2")
	if err != nil {
		t.Error(err)
	}

	if user.Name != "user2" {
		t.Error(err)
	}
}

func TestAllIdAndNames(t *testing.T) {

}

func TestCreate(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(false)

	r := UserRepository{db: db}

	id := "2222"
	name := "BBBB"

	// Mock設定
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("created_at","updated_at","name","login_id","password") VALUES ($1,$2,$3,$4,$5) RETURNING "users"."id"`)).
		WithArgs(Any{}, Any{}, id, name, Any{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	// 実行
	err = r.Create("2222", "BBBB", "dummy")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "id" = $1, "login_id" = $2, "name" = $3, "password" = $4, "updated_at" = $5  WHERE "users"."id" = $6`)).
		WithArgs(2, "change", "user2", sqlmock.AnyArg(), sqlmock.AnyArg(), 2).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	repo := UserRepository{
		db: db,
	}

	err = repo.Update(2, "user2", "change", "password")
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"  WHERE (id=$1)`)).
		WithArgs(2).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	repo := UserRepository{
		db: db,
	}

	err = repo.Delete(2)
	if err != nil {
		t.Error(err)
	}
}

func TestExists(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "users"  WHERE ("users"."name" = $1) LIMIT 1`)).
		WithArgs("user2").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "login_id"}).
			AddRow("1", "user2", "i_am_user1"))

	repo := UserRepository{
		db: db,
	}

	ok, err := repo.Exists(models.User{Name: "user2"})
	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Fatal("false")
	}
}
