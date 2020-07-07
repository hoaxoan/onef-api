package model

import "time"

type Device struct {
	Id      int64     `json:"id,omitempty" bson:"id"`
	OwnerId int64     `json:"owner_id,omitempty" bson:"owner_id"`
	UUId    string    `json:"uuid,omitempty" bson:"uuid"`
	Name    string    `json:"name,omitempty" bson:"name"`
	Created time.Time `json:"created,omitempty" bson:"created"`
	Owner   *User     `json:"owner,omitempty" bson:"owner"`
}

type DeviceRequest struct {
}

type DeviceResponse struct {
	Device  *Device   `json:"device,omitempty" bson:"device"`
	Devices []*Device `json:"devices,omitempty" bson:"devices"`
	Errors  []*Error  `json:"errors,omitempty" bson:"errors"`
}
