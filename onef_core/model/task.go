package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"first_name,omitempty" bsn:"first_name"`
	Description string             `json:"last_name,omitempty" bson:"las_name"`
	DueDate     string             `json:"phone,omitempty" bson:"phone"`
	IsCompleted string             `json:"email,omitempty" bson:"email"`
	IsFlaged    string             `json:"password,omitempty" bson:"password"`
}

type TaskRequest struct {
}

type TaskResponse struct {
	Task   *Task    `json:"task,omitempty" bson:"task"`
	Tasks  []*Task  `json:"tasks,omitempty" bson:"tasks"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}
