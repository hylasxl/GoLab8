package main

import (
	"golang_test_day_07/database"
	"golang_test_day_07/routes"
	"log"
)

func main() {
	database.Connect()
	router := routes.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}

}
