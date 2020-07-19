package model

type CirclePost struct {
	Model
	CircleId *int64 `json:"cirlce_id,omitempty" gorm:"column:cirlce_id" bson:"cirlce_id"`
	PostId   *int64 `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
}

type CirclePostRequest struct {
}

type CirclePostResponse struct {
	CirclePost  *CirclePost  `json:"circle_post,omitempty" bson:"circle_post"`
	CirclePosts []CirclePost `json:"circle_posts,omitempty" bson:"circle_posts"`
	Errors      []*Error     `json:"errors,omitempty" bson:"errors"`
}
