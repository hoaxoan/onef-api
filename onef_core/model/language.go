package model

import "time"

type Language struct {
	Id      int64     `json:"id,omitempty" bson:"id"`
	Code    string    `json:"code,omitempty" bsn:"code"`
	Name    string    `json:"name,omitempty" bson:"name"`
	Created time.Time `json:"created,omitempty" bson:"created"`
}

type LanguageRequest struct {
}

type LanguageResponse struct {
	Language  *Language   `json:"language,omitempty" bson:"language"`
	Languages []*Language `json:"languages,omitempty" bson:"languages"`
	Errors    []*Error    `json:"errors,omitempty" bson:"errors"`
}
