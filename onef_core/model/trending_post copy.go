package model

import "time"

type TrendingPost struct {
	Id      int       `json:"id,omitempty" bson:"id"`
	PostId  int64     `json:"post_id,omitempty" bson:"post_id"`
	Created time.Time `json:"created,omitempty" bson:"created"`
	Post    *Post     `json:"post,omitempty" bson:"post"`
}

type TrendingPostRequest struct {
}

type TrendingPostResponse struct {
	TrendingPost  *TrendingPost   `json:"trending_post,omitempty" bson:"trending_post"`
	TrendingPosts []*TrendingPost `json:"trending_posts,omitempty" bson:"trending_posts"`
	Errors        []*Error        `json:"errors,omitempty" bson:"errors"`
}
