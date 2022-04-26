package uris

import (
	"log"
	"strconv"
)

type UserUri struct {
	Id string `uri:"id"`
}

func (uri *UserUri) GetUserId() int {
	return uri.ToInt()
}

func (uri *UserUri) ToInt() int {
	log.Println(uri.Id)
	id, err := strconv.Atoi(uri.Id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}
