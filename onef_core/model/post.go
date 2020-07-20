package model

import "time"

type Post struct {
	Model
	UUId            string        `json:"uuid,omitempty" gorm:"column:uuid" bson:"uuid"`
	Text            string        `json:"text,omitempty" gorm:"column:text" bson:"text"`
	CommentsEnabled bool          `json:"comments_enabled,omitempty" gorm:"column:comments_enabled" bson:"comments_enabled"`
	PublicReactions bool          `json:"public_reactions,omitempty" gorm:"column:public_reactions" bson:"public_reactions"`
	CreatorId       *int64        `json:"creator_id,omitempty" gorm:"column:creator_id" bson:"creator_id"`
	CommunityId     *int64        `json:"community_id,omitempty" gorm:"column:community_id" bson:"community_id"`
	LanguageId      *int64        `json:"language_id,omitempty" gorm:"column:language_id" bson:"language_id"`
	MediaHeight     float32       `json:"media_height,omitempty" gorm:"column:media_height" bson:"media_height"`
	MediaWidth      float32       `json:"media_width,omitempty" gorm:"column:media_width" bson:"media_width"`
	MediaThumbnail  string        `json:"media_thumbnail,omitempty" gorm:"column:media_thumbnail" bson:"media_thumbnail"`
	Status          string        `json:"status,omitempty" gorm:"column:status" bson:"status"`
	Created         time.Time     `json:"created,omitempty" gorm:"column:created" bson:"created"`
	Modified        time.Time     `json:"modified,omitempty" gorm:"column:modified" bson:"modified"`
	IsMuted         bool          `json:"is_muted,omitempty" gorm:"_" sql:"-" bson:"is_muted"`
	IsEncircled     bool          `json:"is_encircled,omitempty" gorm:"_" sql:"-" bson:"is_encircled"`
	IsEdited        bool          `json:"is_edited,omitempty" gorm:"column:is_edited" bson:"is_edited"`
	IsClosed        bool          `json:"is_closed,omitempty" gorm:"column:is_closed" bson:"is_closed"`
	IsDeleted       bool          `json:"is_deleted,omitempty" gorm:"column:is_deleted" bson:"is_deleted"`
	IsReported      bool          `json:"is_reported,omitempty" gorm:"_" sql:"-" bson:"is_reported"`
	Creator         *User         `json:"creator,omitempty" gorm:"foreignkey:CreatorId" bson:"creator"`
	Community       *Community    `json:"community,omitempty" gorm:"foreignkey:CommunityId" bson:"community"`
	Language        *Language     `json:"language,omitempty" gorm:"foreignkey:LanguageId" bson:"language"`
	PostComments    []PostComment `json:"comments,omitempty" gorm:"_" bson:"comments"`
	PostMedias      []PostMedia   `json:"media,omitempty" gorm:"_" bson:"media"`
	PostReaction    *PostReaction `json:"reaction,omitempty" gorm:"_" bson:"reaction"`
}

func (Post) TableName() string {
	return "post"
}

type CreatePostRequest struct {
	UUId          string   `json:"uuid,omitempty"`
	Text          string   `json:"text,omitempty"`
	CommunityName string   `json:"community_name,omitempty"`
	CircleId      string   `json:"circle_id,omitempty"`
	IsDraft       bool     `json:"is_draft,omitempty"`
	Creator       *User    `json:"creator,omitempty"`
	CircleIds     []string `json:"circle_ids,omitempty"`
}

type EditPostRequest struct {
	Text string `json:"text,omitempty"`
}

type PostRequest struct {
	ListIds   []int64 `json:"list_id,omitempty"`
	CircleIds []int64 `json:"circle_id,omitempty"`
	Count     int64   `json:"count,omitempty"`
	MaxId     int64   `json:"max_id,omitempty"`
	UserName  string  `json:"username,omitempty"`
}

type PostResponse struct {
	Post   *Post    `json:"post,omitempty" bson:"post"`
	Posts  []Post   `json:"posts,omitempty" bson:"posts"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}
