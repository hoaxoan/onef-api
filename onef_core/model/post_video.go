package model

type PostVideo struct {
	Model
	PostId          *int64  `json:"post_id,omitempty" gorm:"column:post_id" bson:"post_id"`
	Hash            string  `json:"hash,omitempty" gorm:"column:hash" bson:"hash"`
	Height          float32 `json:"height,omitempty" gorm:"column:height" bson:"height"`
	Width           float32 `json:"width,omitempty" gorm:"column:width" bson:"width"`
	Duration        float32 `json:"duration,omitempty" gorm:"column:duration" bson:"duration"`
	File            string  `json:"file,omitempty" gorm:"column:file" bson:"file"`
	FileFormat      string  `json:"file_format,omitempty" gorm:"column:file_format" bson:"file_format"`
	Thumbnail       string  `json:"thumbnail,omitempty" gorm:"column:thumbnail" bson:"thumbnail"`
	ThumbnailHeight float32 `json:"thumbnail_height,omitempty" gorm:"column:thumbnail_height" bson:"thumbnail_height"`
	ThumbnailWidth  float32 `json:"thumbnail_width,omitempty" gorm:"column:thumbnail_width" bson:"thumbnail_width"`
	Post            *Post   `json:"post,omitempty" gorm:"foreignkey:PostId" bson:"post"`
}

func (PostImage) PostVideo() string {
	return "postvideo"
}

type PostVideoRequest struct {
}

type PostVideoResponse struct {
	PostVideo  *PostVideo  `json:"post_video,omitempty" bson:"post_video"`
	PostVideos []PostVideo `json:"post_videos,omitempty" bson:"post_videos"`
	Errors     []*Error    `json:"errors,omitempty" bson:"errors"`
}
