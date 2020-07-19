package model

import "time"

type Circle struct {
	Model
	Name       string    `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Color      string    `json:"color,omitempty" gorm:"column:color" bson:"color"`
	Created    time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	CreatorId  *int64    `json:"creator_id,omitempty" gorm:"column:creator_id" bson:"creator_id"`
	Creator    *User     `json:"creator,omitempty" gorm:"foreignkey:CreatorId" bson:"creator"`
	UsersCount int64     `json:"users_count" gorm:"column:users_count" bson:"users_count"`
	Users      []User    `json:"users,omitempty" bson:"users"`
}

type CircleRequest struct {
}

type CreateCircleRequest struct {
}

type UpdateCircleRequest struct {
	Id        int64    `json:"id,omitempty" bson:"id"`
	Name      string   `json:"name,omitempty" bson:"name"`
	Color     string   `json:"color,omitempty" bson:"color"`
	UserNames []string `json:"usernames,omitempty" bson:"usernames"`
}

type CircleResponse struct {
	Circle  *Circle  `json:"circle,omitempty" bson:"circle"`
	Circles []Circle `json:"circles,omitempty" bson:"circles"`
	Errors  []*Error `json:"errors,omitempty" bson:"errors"`
}
