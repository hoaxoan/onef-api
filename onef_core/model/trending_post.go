package model

import "time"

type TrendingPost struct {
	Model
	PostId  *int64    `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	Created time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	Post    *Post     `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

func (TrendingPost) TableName() string {
	return "trendingpost"
}

type TrendingPostRequest struct {
	Count int64 `json:"count,omitempty"`
	MaxId int64 `json:"max_id,omitempty"`
	MinId int64 `json:"min_id,omitempty"`
}

type TrendingPostResponse struct {
	TrendingPost  *TrendingPost  `json:"trending_post,omitempty" bson:"trending_post"`
	TrendingPosts []TrendingPost `json:"trending_posts,omitempty" bson:"trending_posts"`
	Errors        []*Error       `json:"errors,omitempty" bson:"errors"`
}
