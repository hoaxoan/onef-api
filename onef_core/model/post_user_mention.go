package model

type PostUserMention struct {
	Model
	UserId *int64 `json:"user_id,omitempty" gorm:"column:user_id" bson:"user_id"`
	PostId *int64 `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	User   *User  `json:"user,omitempty" gorm:"foreignkey:UserId" bson:"user"`
	Post   *Post  `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

type PostUserMentionRequest struct {
}

type PostUserMentionResponse struct {
	PostUserMention  *PostUserMention  `json:"post_user_mention,omitempty" bson:"post_user_mention"`
	PostUserMentions []PostUserMention `json:"post_user_mentions,omitempty" bson:"post_user_mentions"`
	Errors           []*Error          `json:"errors,omitempty" bson:"errors"`
}
