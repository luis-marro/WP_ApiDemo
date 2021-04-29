package model

import (
	"WP_ApiDemo/apiV1/storage"
	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

var DbClientTesting *firestore.Client

func TestMain(m *testing.M) {
	err := godotenv.Load("../../config.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	// Initialize the DB
	DbClientTesting, err = storage.SetUpDatabase("TESTING_CREDENTIALS")
	if err != nil {
		log.Fatal("Error setting DB: ", err)
	}
	defer DbClientTesting.Close()
	log.Println("DB Connection was Successful")
}
