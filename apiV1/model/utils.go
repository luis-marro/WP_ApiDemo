package model

import (
	"cloud.google.com/go/firestore"
	"context"
)

var DbClient *firestore.Client
var Ctx = context.Background()
