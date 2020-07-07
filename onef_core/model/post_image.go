package model

type PostImage struct {
	Id        int     `json:"id,omitempty" bson:"id"`
	PostId    int64   `json:"post_id,omitempty" bson:"post_id"`
	Image     string  `json:"image,omitempty" bson:"image"`
	Height    float32 `json:"height,omitempty" bson:"height"`
	Width     float32 `json:"width,omitempty" bson:"width"`
	Thumbnail string  `json:"thumbnail,omitempty" bson:"thumbnail"`
	Hash      string  `json:"hash,omitempty" bson:"hash"`
	Post      *Post   `json:"post,omitempty" bson:"post"`
}

type PostImageRequest struct {
}

type PostImageResponse struct {
	PostImage  *PostImage   `json:"post_image,omitempty" bson:"post_image"`
	PostImages []*PostImage `json:"post_images,omitempty" bson:"post_images"`
	Errors     []*Error     `json:"errors,omitempty" bson:"errors"`
}
