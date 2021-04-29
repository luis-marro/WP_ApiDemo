package main

import (
	"WP_ApiDemo/apiV1/controllers"
	"WP_ApiDemo/apiV1/model"
	"WP_ApiDemo/apiV1/storage"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/joho/godotenv"
	"log"
)

var DbClient *firestore.Client
var Ctx = context.Background()

func main() {
	// read the env file
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	// Initialize the DB
	DbClient, err := storage.SetUpDatabase("GOOGLE_APPLICATION_CREDENTIALS")
	if err != nil {
		log.Fatal("Error setting DB: ", err)
	}
	defer DbClient.Close()
	log.Println("DB Connection was Successful")
	// Inject the database to the model
	model.DbClient = DbClient

	// initialize the server
	r := controllers.InitServer()
	r.Run(":8080")

}
