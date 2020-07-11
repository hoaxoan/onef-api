package model

import "time"

type Circle struct {
	Model
	Name      string    `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Color     string    `json:"color,omitempty" gorm:"column:color" bson:"color"`
	Created   time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	CreatorId int       `json:"creator_id,omitempty" gorm:"column:creator_id" bson:"creator_id"`
	Creator   *User     `json:"creator,omitempty" gorm:"foreignkey:CreatorId" bson:"creator"`
}

type CircleRequest struct {
}

type CircleResponse struct {
	Circle  *Circle  `json:"circle,omitempty" bson:"circle"`
	Circles []Circle `json:"circles,omitempty" bson:"circles"`
	Errors  []*Error `json:"errors,omitempty" bson:"errors"`
}
