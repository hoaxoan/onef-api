package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mood struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name  string             `json:"name,omitempty" bsn:"name"`
	Color string             `json:"color,omitempty" bson:"color"`
	Code  string             `json:"code,omitempty" bson:"code"`
	Order int                `json:"order,omitempty" bson:"order"`
}

type MoodRequest struct {
}

type MoodResponse struct {
	Mood   *Mood    `json:"mood,omitempty" bson:"mood"`
	Moods  []*Mood  `json:"moods,omitempty" bson:"moods"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}
