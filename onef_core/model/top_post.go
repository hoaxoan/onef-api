package model

import "time"

type TopPost struct {
	Model
	PostId  *int64    `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	Created time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	Post    *Post     `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

func (TopPost) TableName() string {
	return "top_post"
}

type TopPostRequest struct {
	Count                    int64 `json:"count,omitempty"`
	MaxId                    int64 `json:"max_id,omitempty"`
	MinId                    int64 `json:"min_id,omitempty"`
	ExcludeJoinedCommunities bool  `json:"exclude_joined_communities,omitempty"`
}

type TopPostResponse struct {
	TopPost  *TopPost  `json:"top_post,omitempty" bson:"top_post"`
	TopPosts []TopPost `json:"top_posts,omitempty" bson:"top_posts"`
	Errors   []*Error  `json:"errors,omitempty" bson:"errors"`
}
