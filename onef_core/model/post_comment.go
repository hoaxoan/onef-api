package model

import (
	"time"
)

type PostComment struct {
	Id              int          `json:"id,omitempty" bson:"id"`
	PostId          int64        `json:"post_id,omitempty" bson:"post_id"`
	ParentCommentId int64        `json:"parent_comment_id,omitempty" bson:"parent_comment_id"`
	CommenterId     int64        `json:"commenter_id,omitempty" bson:"commenter_id"`
	Text            string       `json:"text,omitempty" bson:"text"`
	LanguageId      int64        `json:"language_id,omitempty" bson:"language_id"`
	IsEdited        bool         `json:"is_edited,omitempty" bson:"is_edited"`
	IsDeleted       bool         `json:"is_deleted,omitempty" bson:"is_deleted"`
	Created         time.Time    `json:"created,omitempty" bson:"created"`
	Modified        time.Time    `json:"modified,omitempty" bson:"modified"`
	Commenter       *User        `json:"commenter,omitempty" bson:"commenter"`
	Post            *Post        `json:"post,omitempty" bson:"post"`
	ParentComment   *PostComment `json:"parent_comment,omitempty" bson:"parent_comment"`
	Language        *Language    `json:"language,omitempty" bson:"language"`
}

type PostCommentRequest struct {
}

type PostCommentResponse struct {
	PostComment  *PostComment   `json:"post_comment,omitempty" bson:"post_comment"`
	PostComments []*PostComment `json:"post_comments,omitempty" bson:"post_comments"`
	Errors       []*Error       `json:"errors,omitempty" bson:"errors"`
}
