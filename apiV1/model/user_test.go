package model

import (
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	testUser := User{
		"Luis Marroquin",
		"luisfer.marroquin1@gmail.com",
		"43294685",
		"Carretera a el Salvador",
		1,
	}
	err := CreateNewUser(testUser)
	if err != nil {
		t.Error("Error inserting the new user: ", err)
	}
}

func TestFetchUser(t *testing.T) {
	_, err := FetchUser("luisfer.marroquin1@gmail.com")
	if err != nil {
		t.Error("Expected user was not found, ", err)
	}

}
