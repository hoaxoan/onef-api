package model

import "time"

type PostReaction struct {
	Model
	ReactorId *int64    `json:"reactor_id,omitempty" gorm:"column:reactor_id" bson:"reactor_id"`
	EmojiId   *int64    `json:"emoji_id,omitempty" gorm:"column:emoji_id" bson:"emoji_id"`
	PostId    *int64    `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	Modified  time.Time `json:"modified,omitempty" gorm:"column:modified" bson:"modified"`
	Reactor   *User     `json:"reactor,omitempty" gorm:"foreignkey:ReactorId" bson:"reactor"`
	Emoji     *Emoji    `json:"ermoji,omitempty" gorm:"foreignkey:EmojiId" bson:"ermoji"`
	Post      *Post     `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

type PostReactionRequest struct {
}

type PostReactionResponse struct {
	PostReaction  *PostReaction  `json:"post_reaction,omitempty" bson:"post_reaction"`
	PostReactions []PostReaction `json:"post_reactions,omitempty" bson:"post_reactions"`
	Errors        []*Error       `json:"errors,omitempty" bson:"errors"`
}
