package model

import (
	"time"
)

type Story struct {
	Id          int         `json:"id,omitempty" bson:"id"`
	Title       string      `json:"title,omitempty" bsn:"title"`
	Description string      `json:"description,omitempty" bson:"description"`
	Note        string      `json:"note,omitempty" bson:"note"`
	Avatar      string      `json:"avatar,omitempty" bson:"avatar"`
	IsFavorite  string      `json:"is_favorite,omitempty" bson:"is_favorite"`
	Category    *Category   `json:"category,omitempty" bson:"category"`
	Mood        *Mood       `json:"mood,omitempty" bson:"mood"`
	ColorRange  *ColorRange `json:"color_range,omitempty" bson:"color_range"`
	Owner       *User       `json:"owner,omitempty" bson:"owner"`
	Created     time.Time   `json:"created,omitempty" bson:"created"`
	Modified    time.Time   `json:"modified,omitempty" bson:"modified"`
}

type StoryRequest struct {
}

type StoryResponse struct {
	Story   *Story   `json:"story,omitempty" bson:"story"`
	Stories []*Story `json:"stories,omitempty" bson:"stories"`
	Errors  []*Error `json:"errors,omitempty" bson:"errors"`
}
