package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tngranados/tech-challenge-time/config"
	"github.com/tngranados/tech-challenge-time/controllers"
	"github.com/tngranados/tech-challenge-time/storage"
)

func main() {
	// Get configuration parameters
	c, err := config.Get()
	if err != nil {
		log.Fatalf("Error getting the configuration paramaters: %v", err)
	}

	// Get database
	db, err := storage.NewDatabase(c.DatabasePath)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	router, err := controllers.SetupRouter(c.Debug, db)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(fmt.Sprintf(":%v", c.APIPort), router)
	if err != nil {
		log.Fatal(err)
	}
}
