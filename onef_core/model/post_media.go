package model

type PostMedia struct {
	Id            int    `json:"id,omitempty" bson:"id"`
	PostId        int64  `json:"post_id,omitempty" bson:"post_id"`
	Type          string `json:"type,omitempty" bson:"type"`
	ContentType   string `json:"content_type,omitempty" bson:"content_type"`
	ObjectId      string `json:"object_id,omitempty" bson:"object_id"`
	ContentObject string `json:"content_object,omitempty" bson:"content_object"`
	Post          *Post  `json:"post,omitempty" bson:"post"`
}

type PostMediaRequest struct {
}

type PostMediaResponse struct {
	PostMedia  *PostMedia   `json:"post_media,omitempty" bson:"post_media"`
	PostMedias []*PostMedia `json:"post_medias,omitempty" bson:"post_medias"`
	Errors     []*Error     `json:"errors,omitempty" bson:"errors"`
}
