package model

import "time"

type Category struct {
	Model
	Name        string    `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Title       string    `json:"title,omitempty"  gorm:"column:title" bson:"title"`
	Description string    `json:"description,omitempty" gorm:"column:description" bson:"description"`
	Avatar      string    `json:"avatar,omitempty" gorm:"column:avatar" bson:"avatar"`
	Cover       string    `json:"cover,omitempty" gorm:"column:cover" bson:"cover"`
	Color       string    `json:"color,omitempty" gorm:"column:color" bson:"color"`
	Order       int32     `json:"order,omitempty" gorm:"column:order" bson:"order"`
	Created     time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	CreatorId   int       `json:"creator_id,omitempty" gorm:"column:creator_id" bson:"creator_id"`
	Creator     *User     `json:"creator,omitempty" gorm:"foreignkey:CreatorId" bson:"creator"`
}

type CategoryRequest struct {
}

type CategoryResponse struct {
	Category   *Category  `json:"category,omitempty" bson:"category"`
	Categories []Category `json:"categories,omitempty" bson:"categories"`
	Errors     []*Error   `json:"errors,omitempty" bson:"errors"`
}
