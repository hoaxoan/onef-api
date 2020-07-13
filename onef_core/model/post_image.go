package model

type PostImage struct {
	Model
	PostId    *int64  `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	Image     string  `json:"image,omitempty" gorm:"column:image" bson:"image"`
	Height    float32 `json:"height,omitempty" gorm:"column:height" bson:"height"`
	Width     float32 `json:"width,omitempty" gorm:"column:width" bson:"width"`
	Thumbnail string  `json:"thumbnail,omitempty" gorm:"column:thumbnail" bson:"thumbnail"`
	Hash      string  `json:"hash,omitempty" gorm:"column:hash" bson:"hash"`
	Post      *Post   `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

func (PostImage) TableName() string {
	return "postimage"
}

type PostImageRequest struct {
}

type PostImageResponse struct {
	PostImage  *PostImage  `json:"post_image,omitempty" bson:"post_image"`
	PostImages []PostImage `json:"post_images,omitempty" bson:"post_images"`
	Errors     []*Error    `json:"errors,omitempty" bson:"errors"`
}
