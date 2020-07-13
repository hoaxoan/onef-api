package model

import (
	"time"
)

type PostComment struct {
	Model
	PostId          int64        `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	ParentCommentId int64        `json:"parent_comment_id,omitempty" gorm:"column:parent_comment_id" bson:"parent_comment_id"`
	CommenterId     int64        `json:"commenter_id,omitempty" gorm:"column:commenter_id" bson:"commenter_id"`
	Text            string       `json:"text,omitempty" gorm:"column:text" bson:"text"`
	LanguageId      int64        `json:"language_id,omitempty" gorm:"column:language_id" bson:"language_id"`
	IsEdited        bool         `json:"is_edited,omitempty" gorm:"column:is_edited" bson:"is_edited"`
	IsDeleted       bool         `json:"is_deleted,omitempty" gorm:"column:is_deleted" bson:"is_deleted"`
	Created         time.Time    `json:"created,omitempty" gorm:"column:created" bson:"created"`
	Modified        time.Time    `json:"modified,omitempty" gorm:"column:modified" bson:"modified"`
	Commenter       *User        `json:"commenter,omitempty" gorm:"foreignkey:CommenterId" bson:"commenter"`
	Post            *Post        `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
	ParentComment   *PostComment `json:"parent_comment,omitempty" gorm:"foreignkey:ParentCommentId" bson:"parent_comment"`
	Language        *Language    `json:"language,omitempty" gorm:"foreignkey:LanguageId" bson:"language"`
}

func (PostComment) TableName() string {
	return "postcomment"
}

type PostCommentRequest struct {
}

type PostCommentResponse struct {
	PostComment  *PostComment  `json:"comment,omitempty" bson:"comment"`
	PostComments []PostComment `json:"comments,omitempty" bson:"comments"`
	Errors       []*Error      `json:"errors,omitempty" bson:"errors"`
}
