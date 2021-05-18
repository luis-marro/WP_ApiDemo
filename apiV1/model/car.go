package model

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
)

// Car represents the model of a car in the Database
type Car struct {
	Make   string `json:"make" firestore:"make" example:"Volvo"`
	Model  string `json:"model" firestore:"model" example:"C30"`
	Engine string `json:"engine" firestore:"engine" example:"2.5T"`
	Year   int    `json:"year" firestore:"year" example:"2013"`
	Id     string `json:"id" example:"9OiBdQ41inBr4KigocWj"`
}

const carCollection = "Car"
const makesCollection = "CarMakes"

// GetAllCars function to get an array with all the car makes in the database
func GetAllCars() ([]string, error) {
	var allCars []string
	query := DbClient.Collection(makesCollection).
		OrderBy("Make", firestore.Desc).
		Documents(Ctx)
	for {
		var car Car
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Error iterating Cars: ", err)
			return nil, err
		}
		err = doc.DataTo(&car)
		if err != nil {
			log.Println("Error on retrieved object: ", err)
			continue
		}
		allCars = append(allCars, car.Make)
	}

	return allCars, nil
}

// GetModelsByMake Function to get the models that belong to a car make
func GetModelsByMake(make string) ([]string, error) {
	var allCars []string
	query := DbClient.Collection(carCollection).
		Where("Make", "==", make).
		OrderBy("Model", firestore.Desc).
		Documents(Ctx)
	for {
		var car Car
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Error iterating Cars: ", err)
			return nil, err
		}
		err = doc.DataTo(&car)
		if err != nil {
			log.Println("Error on retrieved object: ", err)
			continue
		}
		allCars = append(allCars, car.Model)
	}

	return allCars, nil
}

// GetCarEngine Function to lookup in the database the available engines for a pair of
// Make and model.
func GetCarEngine(make, model string) ([]string, error) {
	var allCars []string
	query := DbClient.Collection(carCollection).
		Where("Make", "==", make).
		Where("Model", "==", model).
		Documents(Ctx)
	for {
		var car Car
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Error iterating Cars: ", err)
			return nil, err
		}
		err = doc.DataTo(&car)
		if err != nil {
			log.Println("Error on retrieved object: ", err)
			continue
		}
		allCars = append(allCars, car.Engine)
	}

	return allCars, nil
}

// GetCarReference function to get the database reference for a specific car
func GetCarReference(make, model string) (string, error) {
	query := DbClient.Collection(carCollection).Where("Make", "==", make).
		Where("Model", "==", model).
		Documents(Ctx)
	doc, err := query.Next()
	if err == iterator.Done {
		return "", err
	}
	if err != nil {
		return "", err
	}
	reference := doc.Ref.ID
	return reference, nil
}
