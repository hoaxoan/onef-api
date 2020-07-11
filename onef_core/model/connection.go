package model

type Connection struct {
	Model
	UserId             int         `json:"user_id,omitempty" gorm:"column:user_id" bson:"user_id"`
	TargetUserId       int         `json:"target_user_id,omitempty" gorm:"column:target_user_id" bson:"target_user_id"`
	TargetConnectionId int         `json:"target_connection_id,omitempty" gorm:"column:target_connection_id" bson:"target_connection_id"`
	User               *User       `json:"user,omitempty" gorm:"foreignkey:UserId" bson:"user"`
	TargetUser         *User       `json:"target_user,omitempty" gorm:"foreignkey:TargetUserId" bson:"target_user"`
	TargetConnection   *Connection `json:"target_connection,omitempty" gorm:"foreignkey:TargetConnectionId" bson:"target_connection"`
}

type ConnectionRequest struct {
}

type ConnectionResponse struct {
	Connection  *Connection  `json:"connection,omitempty" bson:"connection"`
	Connections []Connection `json:"connections,omitempty" bson:"connections"`
	Errors      []*Error     `json:"errors,omitempty" bson:"errors"`
}
