package model

type Model struct {
	Id int64 `json:"id,omitempty" gorm:"column:id;primary_key;auto_increment" bson:"id"`
}
