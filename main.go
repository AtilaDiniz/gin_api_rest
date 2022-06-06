package main

import (
	"github.com/AtilaDiniz/api-go-gin/database"
	"github.com/AtilaDiniz/api-go-gin/routes"
)

func main () {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}