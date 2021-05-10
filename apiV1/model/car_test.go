package model

import (
	"fmt"
	"testing"
)

func TestGetAllCars(t *testing.T) {
	allCars, err := GetAllCars()
	if err != nil {
		t.Error("Error retrieving the cars, ", err)
	}
	fmt.Println(allCars)
	if len(allCars) != 4 {
		t.Error("Not all cars were retrieved")
	}
}

func TestGetModelsByMake(t *testing.T) {
	allModels, err := GetModelsByMake("Volvo")
	if err != nil {
		t.Error("Error getting models for a make: ", err)
	}
	fmt.Println(allModels)
}

func TestGetCarEngine(t *testing.T) {
	engines, err := GetCarEngine("Mitsubishi", "Lancer")
	if err != nil {
		t.Error("Error getting the engines from DB, ", err)
	}
	fmt.Println(engines)
}
