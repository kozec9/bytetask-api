package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Setting struct {
	MongoURL     string `json:"mongoUrl"`
	DatabaseName string `json:"databaseName"`
}

type SettingResponse struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	MongoURL     string             `json:"mongoUrl"`
	DatabaseName string             `json:"databaseName"`
}
