package model

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"time"
)

type Part struct {
	Id          string    `json:"id" firestore:"Id"`
	Name        string    `json:"name" firestore:"Name"`
	Description string    `json:"description" firestore:"Description"`
	Keywords    []string  `json:"keywords" firestore:"Keywords"`
	Category    string    `json:"category" firestore:"Category"`
	Price       float32   `json:"price" firestore:"Price"`
	Pictures    []string  `json:"Pictures" firestore:"Pictures"`
	Inventory   int       `json:"Inventory" firestore:"Inventory"`
	IsNew       bool      `json:"IsNew" firestore:"IsNew"`
	CreatedAt   time.Time `json:"CreatedAt" firestore:"CreatedAt"`
	Cars        []string  `json:"fit" firestore:"Fit"`
}

const PartsCollection = "Part"

// GetAllParts Function to get an array of all the parts in the database.
// Returns an array with the parts from the database, and an error in case it happened.
func GetAllParts() ([]Part, error) {
	var allParts []Part
	query := DbClient.Collection(PartsCollection).
		Where("Inventory", ">", 0).
		OrderBy("Inventory", firestore.Desc).
		Documents(Ctx)
	for {
		var part Part
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		ref := doc.Ref.ID
		err = doc.DataTo(&part)
		if err != nil {
			log.Println("Error on retrieved object: ", err)
			continue
		}
		part.Id = ref
		allParts = append(allParts, part)
	}

	return allParts, nil
}

// GetPartByID Function to return an specific part from the Database based on it's ID.
// Returns the retrieved part, and an error when it applies.
func GetPartByID(partID string) (Part, error) {
	var retrievedPart Part
	query, err := DbClient.Collection(PartsCollection).
		Doc(partID).Get(Ctx)
	if err != nil {
		return Part{}, err
	}
	err = query.DataTo(&retrievedPart)
	if err != nil {
		return Part{}, err
	}
	retrievedPart.Id = partID
	return retrievedPart, nil
}

// GetPartByName Function to retrieve the information for all the parts that contain some of the
// Keywords sent in the argument.
func GetPartByName(partName string) ([]Part, error) {
	var matchingParts []Part
	partName = strings.ToUpper(partName)
	keywords := strings.Split(partName, "-")
	log.Println("Splitted Keywords: ", keywords)
	query := DbClient.Collection(PartsCollection).Where("Keywords", "array-contains-any", keywords).
		Documents(Ctx)
	for {
		var currentPart Part
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Error iterating Getting Parts by Name result: ", err)
			return nil, err
		}
		err = doc.DataTo(&currentPart)
		if err != nil {
			continue
		}
		ref := doc.Ref.ID
		currentPart.Id = ref
		matchingParts = append(matchingParts, currentPart)
	}

	return matchingParts, nil
}

// DiminishInventory Function that takes a part ID and diminishes the inventory by one in the Database.
func DiminishInventory(partId string) error {
	_, err := DbClient.Collection(PartsCollection).Doc(partId).Update(Ctx, []firestore.Update{
		{
			Path:  "Inventory",
			Value: firestore.Increment(-1),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateNewPart Function used to create a new Part from scratch in the database.
func CreateNewPart(name, description, category string, price float32, cars, pictures []string,
	isNew bool, inventory int) (string, error) {
	upper := strings.ToUpper(name)
	keywords := strings.Fields(upper)
	var newPart = Part{
		Name:        name,
		Description: description,
		Category:    category,
		Keywords:    keywords,
		Price:       price,
		Pictures:    pictures,
		Inventory:   inventory,
		IsNew:       isNew,
		CreatedAt:   time.Now(),
		Cars:        cars,
	}

	ref, _, err := DbClient.Collection(PartsCollection).Add(Ctx, newPart)
	if err != nil {
		return "", err
	}

	return ref.ID, nil
}

// UpdatePart Function to update the fields of a part in the DB.
// Returns an error if the update operation fails.
func UpdatePart(partId, name, description, category string, price float32, cars, pictures []string,
	isNew bool, inventory int) error {
	upper := strings.ToUpper(name)
	keywords := strings.Fields(upper)
	_, err := DbClient.Collection(PartsCollection).Doc(partId).Update(Ctx, []firestore.Update{
		{
			Path:  "Description",
			Value: description,
		},
		{
			Path:  "Fit",
			Value: cars,
		},
		{
			Path:  "Inventory",
			Value: inventory,
		},
		{
			Path:  "IsNew",
			Value: isNew,
		},
		{
			Path:  "Keywords",
			Value: keywords,
		},
		{
			Path:  "Name",
			Value: name,
		},
		{
			Path:  "Pictures",
			Value: pictures,
		},
		{
			Path:  "Price",
			Value: price,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
