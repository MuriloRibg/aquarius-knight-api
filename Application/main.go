package main

import (
	"aquarius-knight-api/DataBase"
	"aquarius-knight-api/Routes"
)

type App struct {
	route Routes.Route
}

func (a *App) app() {
	a.route.HandleRequests()
}

func main() {
	DataBase.ConectToDataBase()
	p := &App{route: Routes.Route{}}
	p.app()
}
