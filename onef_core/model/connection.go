package model

type Connection struct {
	Id                 int64       `json:"id,omitempty" bson:"id"`
	UserId             int64       `json:"user_id,omitempty" bson:"user_id"`
	TargetUserId       int64       `json:"target_user_id,omitempty" bson:"target_user_id"`
	TargetConnectionId int64       `json:"target_connection_id,omitempty" bson:"target_connection_id"`
	User               *User       `json:"user,omitempty" bson:"user"`
	TargetUser         *User       `json:"target_user,omitempty" bson:"target_user"`
	TargetConnection   *Connection `json:"target_connection,omitempty" bson:"target_connection"`
}

type ConnectionRequest struct {
}

type ConnectionResponse struct {
	Connection  *Connection   `json:"connection,omitempty" bson:"connection"`
	Connections []*Connection `json:"connections,omitempty" bson:"connections"`
	Errors      []*Error      `json:"errors,omitempty" bson:"errors"`
}
