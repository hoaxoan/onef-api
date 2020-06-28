package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name  string             `json:"name,omitempty" bsn:"name"`
	Color string             `json:"color,omitempty" bson:"color"`
	Code  string             `json:"code,omitempty" bson:"code"`
	Order int                `json:"order,omitempty" bson:"order"`
}

type CategoryRequest struct {
}

type CategoryResponse struct {
	Category   *Category   `json:"category,omitempty" bson:"category"`
	Categories []*Category `json:"categories,omitempty" bson:"categories"`
	Errors     []*Error    `json:"errors,omitempty" bson:"errors"`
}
