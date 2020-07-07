package model

import "time"

type Post struct {
	Id              int64      `json:"id,omitempty" bson:"id"`
	UUId            string     `json:"uuid,omitempty" bson:"uuid"`
	Text            string     `json:"text,omitempty" bson:"text"`
	CommentsEnabled bool       `json:"comments_enabled,omitempty" bson:"comments_enabled"`
	PublicReactions bool       `json:"public_reactions,omitempty" bson:"public_reactions"`
	CreatorId       int64      `json:"creator_id,omitempty" bson:"creator_id"`
	CommunityId     int64      `json:"community_id,omitempty" bson:"community_id"`
	LanguageId      int64      `json:"language_id,omitempty" bson:"language_id"`
	IsEdited        bool       `json:"is_edited,omitempty" bson:"is_edited"`
	IsClosed        bool       `json:"is_closed,omitempty" bson:"is_closed"`
	IsDeleted       bool       `json:"is_deleted,omitempty" bson:"is_deleted"`
	MediaHeight     float32    `json:"media_height,omitempty" bson:"media_height"`
	MediaWidth      float32    `json:"media_width,omitempty" bson:"media_width"`
	MediaThumbnail  string     `json:"media_thumbnail,omitempty" bson:"media_thumbnail"`
	Status          string     `json:"status,omitempty" bson:"status"`
	Created         time.Time  `json:"created,omitempty" bson:"created"`
	Modified        time.Time  `json:"modified,omitempty" bson:"modified"`
	Creator         *User      `json:"creator,omitempty" bson:"creator"`
	Community       *Community `json:"community,omitempty" bson:"community"`
	Language        *Language  `json:"language,omitempty" bson:"language"`
}

type PostRequest struct {
}

type PostResponse struct {
	Post   *Post    `json:"post,omitempty" bson:"post"`
	Posts  []*Post  `json:"posts,omitempty" bson:"posts"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}
