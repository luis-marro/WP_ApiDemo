package model

import (
	"testing"
)

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
	returnedPart, err := GetPartByID("lsibO1uDWYouLi1XQS30")
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
	_, err := GetPartByName("filtro de aire")
	if err != nil {
		t.Error("Failed to fetch the data from the Database: ", err)
	}

}
