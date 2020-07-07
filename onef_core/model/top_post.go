package model

import "time"

type TopPost struct {
	Id      int       `json:"id,omitempty" bson:"id"`
	PostId  int64     `json:"post_id,omitempty" bson:"post_id"`
	Created time.Time `json:"created,omitempty" bson:"created"`
	Post    *Post     `json:"post,omitempty" bson:"post"`
}

type TopPostRequest struct {
}

type TopPostResponse struct {
	TopPost  *TopPost   `json:"top_post,omitempty" bson:"top_post"`
	TopPosts []*TopPost `json:"top_posts,omitempty" bson:"top_posts"`
	Errors   []*Error   `json:"errors,omitempty" bson:"errors"`
}
