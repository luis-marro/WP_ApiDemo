package model

import (
	"WP_ApiDemo/apiV1/storage"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../config.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	// Initialize the DB
	DbClient, err = storage.SetUpDatabase("TESTING_CREDENTIALS")
	if err != nil {
		log.Fatal("Error setting DB: ", err)
	}
	defer DbClient.Close()
	log.Println("DB Connection was Successful")

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetAllParts(t *testing.T) {
	returnedParts, err := GetAllParts()
	if err != nil {
		t.Error("Failed to retrieve parts from the database: ", err)
	}
	if returnedParts == nil {
		t.Error("No objects retrieved, null value")
	}
	if len(returnedParts) == 0 {
		t.Error("The returned parts are not as expected")
	}
	// print the result
	t.Log(returnedParts)
}

func TestGetPartByID(t *testing.T) {
	returnedPart, err := GetPartByID("ptjwTrF4QitIiKFpqN6z")
	// if there is an error, test failed
	if err != nil {
		t.Error("Failed to get part by ID, ", err)
	}
	t.Log(returnedPart)
	// test something that won't return anything
	_, err = GetPartByID("Chojojoy")
	if err == nil {
		t.Error("Test failed, should be empty: ", err)
	}
}

func TestGetPartByName(t *testing.T) {
	parts, err := GetPartByName("filtro de aire agua")
	if err != nil {
		t.Error("Failed to fetch the data from the Database: ", err)
	}
	t.Log(parts)
}

func TestCreateNewPart(t *testing.T) {
	ref, err := CreateNewPart("Bomba de Agua", "Bomba de Agua Aisin WPV800",
		"7FBYk6tt2f6QfW7itCDM", 578.56,
		[]string{"https://www.fcpeuro.com/public/assets/products/173597/large/open-uri20141021-20252-blvd6y.?1496444488"},
		true, 5, "Volvo", "S40 T5")

	if err != nil {
		t.Error("Error inserting the part: ", err)
	}
	t.Log("Doc Ref: ", ref)
}

func TestDiminishInventory(t *testing.T) {
	err := DiminishInventory("cTsgwaWig5SAlkxWr3pQ")
	if err != nil {
		t.Error("Error diminishing the inventory for the item: ", err)
	}
	err = DiminishInventory("Chojojoy")
	if err == nil {
		t.Error("Should be an error, none found")
	}
	t.Log("Success")
}

func TestUpdatePart(t *testing.T) {
	err := UpdatePart("cTsgwaWig5SAlkxWr3pQ", "Filtro de Aceite de Motor", "Filtro de aceite de cartucho",
		"BxszjQokw8o6kqLqnhGF", 88.86, []string{"Au4t843YekUuX5ncRWX1", "FvpxWxT0m4corODgqKJR"},
		[]string{"https://www.fcpeuro.com/public/assets/products/173597/large/open-uri20141021-20252-blvd6y.?1496444488"},
		true, 25)

	if err != nil {
		t.Error("Error updating the part, ", err)
	}
}
