package repository

import "url_manager/model"

var links = make([]model.Link, 0)
var idCnt = 1

func CreateLink(link *model.Link) error {
	links = append(links, model.Link{
		Id:          idCnt,
		LoginId:     link.LoginId,
		Title:       link.Title,
		URL:         link.URL,
		Description: link.Description,
	})
	idCnt++
	return nil
}

func UpdateLink(link *model.Link) error {
	i := FindByFunc(links, func(l model.Link) bool {
		return l.Id == link.Id
	})
	links[i].Title = link.Title
	links[i].URL = link.Title
	links[i].Description = link.Description

	return nil
}

func LinkIndex(loginId string) []model.Link {
	index := make([]model.Link, 0)
	for _, l := range links {
		if l.LoginId == loginId {
			index = append(index, l)
		}
	}
	return index
}

func DeleteLink(id int, loginid string) error {
	i := FindByFunc(links, func(l model.Link) bool {
		return l.Id == id && l.LoginId == loginid
	})
	links = Remove(links, i)
	return nil
}
