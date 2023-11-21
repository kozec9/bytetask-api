package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	MongoURL     string `json:"mongoUrl"`
	DatabaseName string `json:"databaseName"`
}

type ConfigResponse struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	MongoURL     string             `json:"mongoUrl"`
	DatabaseName string             `json:"databaseName"`
}
