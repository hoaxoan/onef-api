package model

import (
	"time"
)

type PostCommentReaction struct {
	Model
	ReactorId     int64        `json:"reactor_id,omitempty" gorm:"column:post_id" bson:"reactor_id"`
	EmojiId       int64        `json:"emoji_id,omitempty" gorm:"column:post_id" bson:"emoji_id"`
	PostCommentId int64        `json:"post_comment_id,omitempty" gorm:"column:post_id" bson:"post_comment_id"`
	Modified      time.Time    `json:"modified,omitempty" gorm:"column:post_id" bson:"modified"`
	Reactor       *User        `json:"reactor,omitempty" gorm:"foreignkey:ReactorId" bson:"reactor"`
	Emoji         *Emoji       `json:"ermoji,omitempty" gorm:"foreignkey:EmojiId" bson:"ermoji"`
	PostComment   *PostComment `json:"post_comment,omitempty" gorm:"foreignkey:PostCommentId" bson:"post_comment"`
}

type PostCommentReactionRequest struct {
}

type PostCommentReactionResponse struct {
	PostCommentReaction  *PostCommentReaction  `json:"post_comment_reaction,omitempty" bson:"post_comment_reaction"`
	PostCommentReactions []PostCommentReaction `json:"post_comment_reactions,omitempty" bson:"post_comment_reaction"`
	Errors               []*Error              `json:"errors,omitempty" bson:"errors"`
}
