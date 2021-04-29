package model

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"time"
)

type Part struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Pictures    []string  `json:"Pictures"`
	Inventory   int       `json:"Inventory"`
	IsNew       bool      `json:"IsNew"`
	CreatedAt   time.Time `json:"CreatedAt"`
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
