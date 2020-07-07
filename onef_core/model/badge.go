package model

import "time"

type Badge struct {
	Id                 int64     `json:"id,omitempty" bson:"id"`
	Keyword            string    `json:"keyword,omitempty" bsn:"keyword"`
	KeywordDescription string    `json:"keyword_description,omitempty" bson:"keyword_description"`
	Created            time.Time `json:"created,omitempty" bson:"created"`
}

type BadgeRequest struct {
}

type BadgeResponse struct {
	Badge  *Badge   `json:"badge,omitempty" bson:"badge"`
	Badges []*Badge `json:"badges,omitempty" bson:"badges"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}
