package model

type StoryCategory struct {
	Id    int    `json:"id,omitempty" bson:"id"`
	Name  string `json:"name,omitempty" bsn:"name"`
	Color string `json:"color,omitempty" bson:"color"`
	Code  string `json:"code,omitempty" bson:"code"`
	Order int    `json:"order,omitempty" bson:"order"`
}

type StoryCategoryRequest struct {
}

type StoryCategoryResponse struct {
	StoryCategory   *StoryCategory   `json:"storycategory,omitempty" bson:"storycategory"`
	StoryCategories []*StoryCategory `json:"storycategories,omitempty" bson:"storycategories"`
	Errors          []*Error         `json:"errors,omitempty" bson:"errors"`
}
