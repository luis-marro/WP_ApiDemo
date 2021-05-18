package main

import (
	"WP_ApiDemo/apiV1/controllers"
	"WP_ApiDemo/apiV1/model"
	"WP_ApiDemo/apiV1/storage"
	_ "WP_ApiDemo/docs"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/joho/godotenv"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

// @title Car Spare Parts API
// @version 1.0
// @description API to run the operations for an ecommerce application for care spare parts
// @host localhost:8080, https://carssparepartsstore.appspot.com
// @BasePath /api/v1

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

	// Swagger Definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

}
