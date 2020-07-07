package model

type PostVideo struct {
	Id              int     `json:"id,omitempty" bson:"id"`
	PostId          int64   `json:"post_id,omitempty" bson:"post_id"`
	Hash            string  `json:"hash,omitempty" bson:"hash"`
	Height          float32 `json:"height,omitempty" bson:"height"`
	Width           float32 `json:"width,omitempty" bson:"width"`
	Duration        float32 `json:"duration,omitempty" bson:"duration"`
	File            string  `json:"file,omitempty" bson:"file"`
	FileFormat      string  `json:"file_format,omitempty" bson:"file_format"`
	Thumbnail       string  `json:"thumbnail,omitempty" bson:"thumbnail"`
	ThumbnailHeight float32 `json:"thumbnail_height,omitempty" bson:"thumbnail_height"`
	ThumbnailWidth  float32 `json:"thumbnail_width,omitempty" bson:"thumbnail_width"`
	Post            *Post   `json:"post,omitempty" bson:"post"`
}

type PostVideoRequest struct {
}

type PostVideoResponse struct {
	PostVideo  *PostVideo   `json:"post_video,omitempty" bson:"post_video"`
	PostVideos []*PostVideo `json:"post_videos,omitempty" bson:"post_videos"`
	Errors     []*Error     `json:"errors,omitempty" bson:"errors"`
}
