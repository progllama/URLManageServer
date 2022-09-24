package repository

import "url_manager/model"

var folders = make([]model.Folder, 0)
var folderIdCounter = 1

func IndexFolder(loginId string) []model.Folder {
	index := make([]model.Folder, 0)
	for _, f := range folders {
		if f.LoginId == loginId {
			index = append(index, f)
		}
	}
	return index
}

func CreateFolder(loginId, title string) error {
	folders = append(folders, model.Folder{
		Id:      folderIdCounter,
		LoginId: loginId,
		Title:   title,
	})
	folderIdCounter++
	return nil
}

func UpdateFolder(id int, title string) {
	i := FindByFunc(folders, func(f model.Folder) bool {
		return f.Id == id
	})
	folders[i] = model.Folder{
		Id:      folders[i].Id,
		LoginId: folders[i].LoginId,
		Title:   title,
	}
}

func DeleteFolder(loginId string, id int) {
	i := FindByFunc(folders, func(f model.Folder) bool {
		return f.Id == id
	})
	folders = Remove(folders, i)
}
