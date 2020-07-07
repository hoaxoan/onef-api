package model

import "time"

type Circle struct {
	Id        int64     `json:"id,omitempty" bson:"id"`
	Name      string    `json:"name,omitempty" bson:"name"`
	Color     string    `json:"color,omitempty" bson:"color"`
	Created   time.Time `json:"created,omitempty" bson:"created"`
	CreatorId int64     `json:"creator_id,omitempty" bson:"creator_id"`
	Creator   *User     `json:"creator,omitempty" bson:"creator"`
}

type CircleRequest struct {
}

type CircleResponse struct {
	Circle  *Circle   `json:"circle,omitempty" bson:"circle"`
	Circles []*Circle `json:"circles,omitempty" bson:"circles"`
	Errors  []*Error  `json:"errors,omitempty" bson:"errors"`
}
