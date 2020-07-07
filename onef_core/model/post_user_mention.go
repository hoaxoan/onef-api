package model

type PostUserMention struct {
	Id     int   `json:"id,omitempty" bson:"id"`
	UserId int64 `json:"user_id,omitempty" bson:"user_id"`
	PostId int64 `json:"post_id,omitempty" bson:"post_id"`
	User   *User `json:"user,omitempty" bson:"user"`
	Post   *Post `json:"post,omitempty" bson:"post"`
}

type PostUserMentionRequest struct {
}

type PostUserMentionResponse struct {
	PostUserMention  *PostUserMention   `json:"post_user_mention,omitempty" bson:"post_user_mention"`
	PostUserMentions []*PostUserMention `json:"post_user_mentions,omitempty" bson:"post_user_mentions"`
	Errors           []*Error           `json:"errors,omitempty" bson:"errors"`
}
