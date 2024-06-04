package graph

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DbClient *mongo.Client
}

type Time struct {
	Time time.Time
}
