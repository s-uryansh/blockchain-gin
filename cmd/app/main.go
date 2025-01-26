package main

import (
	"log"

	"github.com/s-uryansh/blockchain-gin/db"
	"github.com/s-uryansh/blockchain-gin/routes"
)

func main() {

	db.InitDatabase() //connecting to database

	//Setting up router for API, and creating routes
	r := routes.SetupRouter()

	log.Println("Starting server on: 8080")
	r.Run(":8080")

}
