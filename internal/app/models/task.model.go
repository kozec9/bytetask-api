package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	// ID primitive.ObjectID `json:"id" bson:"_id"`
	// InstanceID     primitive.ObjectID `json:"instanceId" bson:"instanceId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	// AssigneeUserID primitive.ObjectID `json:"assigneeUserId" bson:"assigneeUserId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskResponse struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	// InstanceID     primitive.ObjectID `json:"instanceId" bson:"instanceId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	// AssigneeUserID primitive.ObjectID `json:"assigneeUserId" bson:"assigneeUserId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
