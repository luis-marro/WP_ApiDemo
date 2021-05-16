package model

import (
	"google.golang.org/api/iterator"
	"log"
)

type User struct {
	Name           string `json:"name" firestore:"name"`
	Email          string `json:"email" firestore:"email"`
	UnsafePassword string `json:"unsafePassword", firestore:"unsafePassword"`
	SafePassword   string `json:"safePassword" firestore:"safePassword"`
	PhoneNum       string `json:"phoneNumber" firestore:"phoneNumber"`
	Address        string `json:"address" firestore:"address"`
	Payment        int    `json:"payment" firestore:"paymentMethod"`
}

const usersCollection = "User"
const PasswordSalt = "f7e54be9a37e316c9a0cea29b607173b4965fb557ad3fb2a070851052b18c05f"

// CreateNewUser function to create a new user in the Database
func CreateNewUser(newUser User) error {
	_, _, err := DbClient.Collection(usersCollection).Add(Ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

// fetchUser function to get a user from the database using the email as ID
func FetchUser(email string, password string) (User, error) {
	var retrievedUser User
	query := DbClient.Collection(usersCollection).
		Where("email", "==", email).
		Where("safePassword", "==", password).
		Documents(Ctx)
	doc, err := query.Next()
	if err == iterator.Done {
		log.Println("User with Email ", email, " was not found")
		return User{}, err
	}
	if err != nil {
		return User{}, err
	}
	err = doc.DataTo(&retrievedUser)
	if err != nil {
		return User{}, err
	}

	log.Println("Successfully found the user ", retrievedUser)
	return retrievedUser, nil
}
