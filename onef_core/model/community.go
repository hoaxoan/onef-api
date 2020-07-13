package model

import "time"

type Community struct {
	Model
	Name           string    `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Title          string    `json:"title,omitempty" gorm:"column:title" bson:"title"`
	Description    string    `json:"description,omitempty" gorm:"column:description" bson:"description"`
	Rules          string    `json:"rules,omitempty" gorm:"column:rules" bson:"rules"`
	Avatar         string    `json:"avatar,omitempty" gorm:"column:avatar" bson:"avatar"`
	Cover          string    `json:"cover,omitempty" gorm:"column:cover" bson:"cover"`
	Color          string    `json:"color,omitempty" gorm:"column:color" bson:"color"`
	Type           string    `json:"type,omitempty" gorm:"column:type" bson:"type"`
	UserAdjective  string    `json:"user_adjective,omitempty" gorm:"column:user_adjective" bson:"user_adjective"`
	UsersAdjective string    `json:"users_adjective,omitempty" gorm:"column:users_adjective" bson:"users_adjective"`
	InvitesEnabled bool      `json:"invites_enabled,omitempty" gorm:"column:invites_enabled" bson:"invites_enabled"`
	Order          int32     `json:"order,omitempty" gorm:"column:order" bson:"order"`
	Created        time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	CreatorId      int64     `json:"creator_id,omitempty" gorm:"column:creator_id" bson:"creator_id"`
	Creator        *User     `json:"creator,omitempty" gorm:"foreignkey:CreatorId" bson:"creator"`
}

type CommunityRequest struct {
}

type CommunityResponse struct {
	Community   *Community  `json:"community,omitempty" bson:"community"`
	Communities []Community `json:"communities,omitempty" bson:"communities"`
	Errors      []*Error    `json:"errors,omitempty" bson:"errors"`
}
