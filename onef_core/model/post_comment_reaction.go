package model

import (
	"time"
)

type PostCommentReaction struct {
	Id            int          `json:"id,omitempty" bson:"id"`
	ReactorId     int64        `json:"reactor_id,omitempty" bson:"reactor_id"`
	EmojiId       int64        `json:"emoji_id,omitempty" bson:"emoji_id"`
	PostCommentId int64        `json:"post_comment_id,omitempty" bson:"post_comment_id"`
	Modified      time.Time    `json:"modified,omitempty" bson:"modified"`
	Reactor       *User        `json:"reactor,omitempty" bson:"reactor"`
	Emoji         *Emoji       `json:"ermoji,omitempty" bson:"ermoji"`
	PostComment   *PostComment `json:"post_comment,omitempty" bson:"post_comment"`
}

type PostCommentReactionRequest struct {
}

type PostCommentReactionResponse struct {
	PostCommentReaction  *PostCommentReaction   `json:"post_comment_reaction,omitempty" bson:"post_comment_reaction"`
	PostCommentReactions []*PostCommentReaction `json:"post_comment_reactions,omitempty" bson:"post_comment_reaction"`
	Errors               []*Error               `json:"errors,omitempty" bson:"errors"`
}
