package app

import "log"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	db := NewDatabase()
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := NewServer()
	Routes(server.Engine())
	server.Open()
}
