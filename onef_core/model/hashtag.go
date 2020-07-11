package model

import "time"

type Hashtag struct {
	Model
	Name      string    `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Color     string    `json:"color,omitempty" gorm:"column:color" bson:"color"`
	TextColor string    `json:"text_color,omitempty" gorm:"column:text_color" bson:"text_color"`
	Width     int32     `json:"width,omitempty" gorm:"column:width" bson:"width"`
	Height    int32     `json:"height,omitempty" gorm:"column:height" bson:"height"`
	Image     string    `json:"image,omitempty" gorm:"column:image" bson:"image"`
	EmojiId   int       `json:"emoji_id,omitempty" gorm:"column:emoji_id" bson:"emoji_id"`
	Created   time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	Emoji     *Emoji    `json:"emoji,omitempty" gorm:"foreignkey:EmojiId" bson:"emoji"`
}

type HashtagRequest struct {
}

type HashtagResponse struct {
	Hashtag  *Hashtag  `json:"hashtag,omitempty" bson:"hashtag"`
	Hashtags []Hashtag `json:"hashtags,omitempty" bson:"hashtags"`
	Errors   []*Error  `json:"errors,omitempty" bson:"errors"`
}
