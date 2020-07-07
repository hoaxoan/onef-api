package model

import "time"

type Category struct {
	Id          int64     `json:"id,omitempty" bson:"id"`
	Name        string    `json:"name,omitempty" bson:"name"`
	Title       string    `json:"title,omitempty" bson:"title"`
	Description string    `json:"description,omitempty" bson:"description"`
	Avatar      string    `json:"avatar,omitempty" bson:"avatar"`
	Cover       string    `json:"cover,omitempty" bson:"cover"`
	Color       string    `json:"color,omitempty" bson:"color"`
	Code        string    `json:"code,omitempty" bson:"code"`
	Order       int32     `json:"order,omitempty" bson:"order"`
	Created     time.Time `json:"created,omitempty" bson:"created"`
	CreatorId   int64     `json:"creator_id,omitempty" bson:"creator_id"`
	Creator     *User     `json:"creator,omitempty" bson:"creator"`
}

type CategoryRequest struct {
}

type CategoryResponse struct {
	Category   *Category   `json:"category,omitempty" bson:"category"`
	Categories []*Category `json:"categories,omitempty" bson:"categories"`
	Errors     []*Error    `json:"errors,omitempty" bson:"errors"`
}
