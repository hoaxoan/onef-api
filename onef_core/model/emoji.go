package model

import "time"

type Emoji struct {
	Id      int64     `json:"id,omitempty" bson:"id"`
	Keyword string    `json:"keyword,omitempty" bson:"keyword"`
	Color   string    `json:"color,omitempty" bson:"color"`
	Image   string    `json:"image,omitempty" bson:"image"`
	Order   int32     `json:"order,omitempty" bson:"order"`
	GroupId int32     `json:"group_id,omitempty" bson:"group_id"`
	Created time.Time `json:"created,omitempty" bson:"created"`
}

type EmojiRequest struct {
}

type EmojiResponse struct {
	Emoji  *Emoji   `json:"emoji,omitempty" bson:"emoji"`
	Emojis []*Emoji `json:"emojis,omitempty" bson:"emojis"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}
