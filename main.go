package main

import (
	"log"

	"github.com/dach3r/vecker/db"
	"github.com/dach3r/vecker/handlers"
)

func main() {
	if db.VerifyConnection() == false {
		log.Fatal("Error to connect with DB")
		return
	}
	handlers.Manager()
}
