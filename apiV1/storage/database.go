// Package storage provides the methdos to initialize the Connection to the Database
// in GCP
package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
)

// SetUpDatabase handles the variable declaration to stablish a new connection to Firestore DB
// It returns a firestore Client object to be used in the project.
func SetUpDatabase(environment string) (*firestore.Client, error) {
	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv(environment))
	conf := &firebase.Config{
		ProjectID: projectID,
	}
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
