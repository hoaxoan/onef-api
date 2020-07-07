package model

import "time"

type Hashtag struct {
	Id        int64     `json:"id,omitempty" bson:"id"`
	Name      string    `json:"name,omitempty" bson:"name"`
	Color     string    `json:"color,omitempty" bson:"color"`
	TextColor string    `json:"text_color,omitempty" bson:"text_color"`
	Width     int32     `json:"width,omitempty" bson:"width"`
	Height    int32     `json:"height,omitempty" bson:"height"`
	Image     string    `json:"image,omitempty" bson:"image"`
	EmojiId   int64     `json:"emoji_id,omitempty" bson:"emoji_id"`
	Created   time.Time `json:"created,omitempty" bson:"created"`
	Emoji     *Emoji    `json:"emoji,omitempty" bson:"emoji"`
}

type HashtagRequest struct {
}

type HashtagResponse struct {
	Hashtag  *Hashtag   `json:"hashtag,omitempty" bson:"hashtag"`
	Hashtags []*Hashtag `json:"hashtags,omitempty" bson:"hashtags"`
	Errors   []*Error   `json:"errors,omitempty" bson:"errors"`
}
