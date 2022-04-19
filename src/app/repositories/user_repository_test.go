package repositories

import (
	"testing"
	"url_manager/app/models"
)

func TestGetAll(t *testing.T) {

}

func TestFind(t *testing.T) {

}

func TestUserCreate(t *testing.T) {
	type args struct {
		models.User
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "normal",
			args: args{models.User{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// repo := UserRepository{}
			// if got := repo.Create(tt.args.User); got != tt.want {
			// 	t.Errorf("got = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestUserUpdate(t *testing.T) {

}

func TestUserDelete(t *testing.T) {

}

func TestExists(t *testing.T) {

}
