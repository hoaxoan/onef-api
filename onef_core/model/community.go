package model

import "time"

type Community struct {
	Id             int64     `json:"id,omitempty" bson:"id"`
	Name           string    `json:"name,omitempty" bson:"name"`
	Title          string    `json:"title,omitempty" bson:"title"`
	Description    string    `json:"description,omitempty" bson:"description"`
	Rules          string    `json:"rules,omitempty" bson:"rules"`
	Avatar         string    `json:"avatar,omitempty" bson:"avatar"`
	Cover          string    `json:"cover,omitempty" bson:"cover"`
	Color          string    `json:"color,omitempty" bson:"color"`
	Type           string    `json:"type,omitempty" bson:"type"`
	UserAdjective  string    `json:"user_adjective,omitempty" bson:"user_adjective"`
	UsersAdjective string    `json:"users_adjective,omitempty" bson:"users_adjective"`
	InvitesEnabled bool      `json:"invites_enabled,omitempty" bson:"invites_enabled"`
	Order          int32     `json:"order,omitempty" bson:"order"`
	Created        time.Time `json:"created,omitempty" bson:"created"`
	CreatorId      int64     `json:"creator_id,omitempty" bson:"creator_id"`
	Creator        *User     `json:"creator,omitempty" bson:"creator"`
}

type CommunityRequest struct {
}

type CommunityResponse struct {
	Community   *Community   `json:"community,omitempty" bson:"community"`
	Communities []*Community `json:"communities,omitempty" bson:"communities"`
	Errors      []*Error     `json:"errors,omitempty" bson:"errors"`
}
