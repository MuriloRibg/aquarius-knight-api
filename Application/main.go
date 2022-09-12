package main

import (
	"aquarius-knight-api/DataBase"
	"aquarius-knight-api/Routes"
)

func main() {
	DataBase.ConectToDataBase()
	Routes.HandleRequests()
}
