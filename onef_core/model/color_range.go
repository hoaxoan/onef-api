package model

type ColorRange struct {
	Id    int    `json:"id,omitempty" bson:"id"`
	Name  string `json:"name,omitempty" bsn:"name"`
	Color string `json:"color,omitempty" bson:"color"`
	Start string `json:"start,omitempty" bson:"start"`
	End   string `json:"end,omitempty" bson:"end"`
}

type ColorRangeRequest struct {
}

type ColorRangeResponse struct {
	ColorRange  *ColorRange   `json:"color_range,omitempty" bson:"color_range"`
	ColorRanges []*ColorRange `json:"color_ranges,omitempty" bson:"color_ranges"`
	Errors      []*Error      `json:"errors,omitempty" bson:"errors"`
}
