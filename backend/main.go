package main

import (
	"kredit/backend/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := api.SetupDB()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
