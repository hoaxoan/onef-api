package model

type Story struct {
	Id          int    `json:"id,omitempty" bson:"id"`
	Name        string `json:"first_name,omitempty" bsn:"first_name"`
	Description string `json:"last_name,omitempty" bson:"las_name"`
	DueDate     string `json:"phone,omitempty" bson:"phone"`
	IsCompleted string `json:"email,omitempty" bson:"email"`
	IsFlaged    string `json:"password,omitempty" bson:"password"`
}

type StoryRequest struct {
}

type StoryResponse struct {
	Story   *Story   `json:"story,omitempty" bson:"story"`
	Stories []*Story `json:"stories,omitempty" bson:"stories"`
	Errors  []*Error `json:"errors,omitempty" bson:"errors"`
}
