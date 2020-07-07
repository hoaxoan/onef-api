package model

import "time"

type EmojiGroup struct {
	Id              int64     `json:"id,omitempty" bson:"id"`
	Keyword         string    `json:"keyword,omitempty" bson:"keyword"`
	Color           string    `json:"color,omitempty" bson:"color"`
	Order           int32     `json:"order,omitempty" bson:"order"`
	IsReactionGroup bool      `json:"is_reaction_group,omitempty" bson:"is_reaction_group"`
	Created         time.Time `json:"created,omitempty" bson:"created"`
}

type EmojiGroupRequest struct {
}

type EmojiGroupResponse struct {
	EmojiGroup  *EmojiGroup   `json:"emojigroup,omitempty" bson:"emojigroup"`
	EmojiGroups []*EmojiGroup `json:"emojigroups,omitempty" bson:"emojigroups"`
	Errors      []*Error      `json:"errors,omitempty" bson:"errors"`
}
