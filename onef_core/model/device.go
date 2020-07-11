package model

import "time"

type Device struct {
	Model
	UUId    string    `json:"uuid,omitempty" gorm:"column:uuid" bson:"uuid"`
	Name    string    `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Created time.Time `json:"created,omitempty" gorm:"column:created" bson:"created"`
	OwnerId int       `json:"owner_id,omitempty" gorm:"column:owner_id" bson:"owner_id"`
	Owner   *User     `json:"owner,omitempty" gorm:"foreignkey:OwnerId" bson:"owner"`
}

type DeviceRequest struct {
}

type DeviceResponse struct {
	Device  *Device  `json:"device,omitempty" bson:"device"`
	Devices []Device `json:"devices,omitempty" bson:"devices"`
	Errors  []*Error `json:"errors,omitempty" bson:"errors"`
}
