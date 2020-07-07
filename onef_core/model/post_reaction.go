package model

import "time"

type PostReaction struct {
	Id        int       `json:"id,omitempty" bson:"id"`
	ReactorId int64     `json:"reactor_id,omitempty" bson:"reactor_id"`
	EmojiId   int64     `json:"emoji_id,omitempty" bson:"emoji_id"`
	PostId    int64     `json:"post_id,omitempty" bson:"post_id"`
	Modified  time.Time `json:"modified,omitempty" bson:"modified"`
	Reactor   *User     `json:"reactor,omitempty" bson:"reactor"`
	Emoji     *Emoji    `json:"ermoji,omitempty" bson:"ermoji"`
	Post      *Post     `json:"post,omitempty" bson:"post"`
}

type PostReactionRequest struct {
}

type PostReactionResponse struct {
	PostReaction  *PostReaction   `json:"post_reaction,omitempty" bson:"post_reaction"`
	PostReactions []*PostReaction `json:"post_reactions,omitempty" bson:"post_reactions"`
	Errors        []*Error        `json:"errors,omitempty" bson:"errors"`
}
