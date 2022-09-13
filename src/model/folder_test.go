package model_test

import (
	"encoding/json"
	"testing"
	"url_manager/model"
)

func TestNormalCaseFolderEncode(t *testing.T) {
	folder := model.Folder{Name: "alice"}
	b, err := json.Marshal(folder)
	if err != nil {
		t.Error(err)
	}
	if string(b) != `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"alice","ManagerID":null,"Folders":null,"Links":null,"UserID":0}` {
		t.Errorf("unexpected %s", string(b))
	}
}

func TestNormalCaseFolderDecode(t *testing.T) {
	type args struct {
		json []byte
	}
	tests := []struct {
		name string
		args args
		want model.Folder
	}{
		{
			name: "lower",
			args: args{json: []byte(`{"name": "alice"}`)},
			want: model.Folder{Name: "alice"},
		},
		{
			name: "upper",
			args: args{json: []byte(`{"NAME": "alice"}`)},
			want: model.Folder{Name: "alice"},
		},
		{
			name: "blank",
			args: args{json: []byte("{}")},
			want: model.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got model.Folder
			json.Unmarshal(tt.args.json, &got)
			if tt.want.Name != got.Name {
				t.Errorf("want %s, got %s", tt.want.Name, got.Name)
			}
		})
	}
}

func TestErrorCaseFolderDecode(t *testing.T) {
	t.Run("invalid format", func(t *testing.T) {
		var f model.Folder
		err := json.Unmarshal([]byte(`{name: "alice"}`), &f)
		if err == nil {
			t.Error(err)
		}
		if err.Error() != "invalid character 'n' looking for beginning of object key string" {
			t.Error(err)
		}
	})
}

func TestNormalCaseFolderValidation(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "alphabet", args: args{name: "Alice"}},
		{name: "number", args: args{name: "1234"}},
		{name: "alphabet and number", args: args{name: "alice12345"}},
		{name: "length 1", args: args{name: "1234"}},
		{name: "length 256", args: args{name: getLongName()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := model.Folder{
				Name: tt.args.name,
			}
			if err := valid(f); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestErrorCaseFolderValidation(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "neither alphabetic or numeric",
			args: args{name: "%ri3a;"},
			err:  "Key: 'Folder.Name' Error:Field validation for 'Name' failed on the 'alphanum' tag",
		},
		{
			name: "blank",
			args: args{name: ""},
			err:  "Key: 'Folder.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
		{
			name: "length 257",
			args: args{name: getLongName() + "a"},
			err:  "Key: 'Folder.Name' Error:Field validation for 'Name' failed on the 'max' tag",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := model.Folder{
				Name: tt.args.name,
			}
			err := valid(f)
			if err == nil {
				t.Error("should error")
			}
			if err.Error() != tt.err {
				t.Error(err)
			}
		})
	}
}

func TestNormalCaseFolderGormConstraint(t *testing.T) {
	type args struct {
		folder model.Folder
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "length 1",
			args: args{
				folder: model.Folder{
					Name: "a",
				},
			},
		},
		{
			name: "length 256",
			args: args{
				folder: model.Folder{
					Name: getLongName(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := createTestDB()
			if err != nil {
				t.Fatal(err)
			}
			con, err := db.DB()
			if err != nil {
				t.Fatal(err)
			}
			defer con.Close()
			err = db.AutoMigrate(&model.Folder{})
			if err != nil {
				t.Fatal(err)
			}

			if err := db.Create(&(tt.args.folder)).Error; err != nil {
				t.Error(err)
			}
		})
	}
}

func TestErrorCaseFolderGormConstraint(t *testing.T) {
	type args struct {
		folder model.Folder
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "length 0",
			args: args{
				folder: model.Folder{
					Name: "",
				},
			},
			err: "CHECK constraint failed: chk_folders_name",
		},
		{
			name: "length 257",
			args: args{
				folder: model.Folder{
					Name: getLongName() + "a",
				},
			},
			err: "CHECK constraint failed: chk_folders_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := createTestDB()
			if err != nil {
				t.Fatal(err)
			}
			con, err := db.DB()
			if err != nil {
				t.Fatal(err)
			}
			defer con.Close()
			err = db.AutoMigrate(&model.Folder{})
			if err != nil {
				t.Fatal(err)
			}

			f := tt.args.folder
			err = db.Create(&f).Error
			if err == nil {
				t.Error("should error")
			}

			if err != nil && err.Error() != tt.err {
				t.Error(err)
			}
		})
	}
}
