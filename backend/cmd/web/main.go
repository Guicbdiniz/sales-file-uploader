package main

import (
	"fmt"
	"guicbdiniz/hubla/backend/internal/api"
	"guicbdiniz/hubla/backend/internal/models"
	"log"
	"net/http"
)

const port = 80

func main() {
	models, err := models.CreatePostgresModels()
	if err != nil {
		log.Fatalf("Error captured while creating the models, %v", err)
	}
	defer models.Clear()

	err = models.InitDB()
	if err != nil {
		log.Fatalf("Error captured while initiating the database, %v", err)
	}

	api, err := api.CreateAPI(models)

	if err != nil {
		log.Fatalf("Error captured while creating the API, %v", err)
	}

	log.Printf("Starting API at port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), api))
}
