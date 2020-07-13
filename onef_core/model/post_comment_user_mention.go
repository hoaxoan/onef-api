package model

type PostCommentUserMention struct {
	Model
	UserId        int64        `json:"user_id,omitempty" gorm:"column:post_id" bson:"user_id"`
	PostCommentId int64        `json:"post_comment_id,omitempty" gorm:"column:post_id" bson:"post_comment_id"`
	User          *User        `json:"user,omitempty" gorm:"foreignkey:UserId" bson:"user"`
	PostComment   *PostComment `json:"post_comment,omitempty" gorm:"foreignkey:PostCommentId" bson:"post_comment"`
}

type PostCommentUserMentionRequest struct {
}

type PostCommentUserMentionResponse struct {
	PostCommentUserMention  *PostCommentUserMention  `json:"post_comment_user_mention,omitempty" bson:"post_comment_user_mention"`
	PostCommentUserMentions []PostCommentUserMention `json:"post_comment_user_mentions,omitempty" bson:"post_comment_user_mentions"`
	Errors                  []*Error                 `json:"errors,omitempty" bson:"errors"`
}
