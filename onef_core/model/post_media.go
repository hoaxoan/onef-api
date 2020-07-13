package model

type PostMedia struct {
	Model
	PostId        *int64 `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	Type          string `json:"type,omitempty" gorm:"column:type" bson:"type"`
	ContentType   string `json:"content_type,omitempty" gorm:"column:content_type" bson:"content_type"`
	ObjectId      string `json:"object_id,omitempty" gorm:"column:object_id" bson:"object_id"`
	ContentObject string `json:"content_object,omitempty" gorm:"column:content_object" bson:"content_object"`
	Post          *Post  `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

func (PostImage) PostMedia() string {
	return "postmedia"
}

type PostMediaRequest struct {
}

type PostMediaResponse struct {
	PostMedia  *PostMedia  `json:"post_media,omitempty" bson:"post_media"`
	PostMedias []PostMedia `json:"post_medias,omitempty" bson:"post_medias"`
	Errors     []*Error    `json:"errors,omitempty" bson:"errors"`
}
