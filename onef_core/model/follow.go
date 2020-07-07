package model

type Follow struct {
	Id             int64 `json:"id,omitempty" bson:"id"`
	UserId         int32 `json:"user_id,omitempty" bson:"user_id"`
	FollowedUserId int32 `json:"followed_user_id,omitempty" bson:"followed_user_id"`
	User           *User `json:"user,omitempty" bson:"user"`
	FollowedUser   *User `json:"followed_user,omitempty" bson:"followed_user"`
}

type FollowsRequest struct {
}

type FollowResponse struct {
	Follow  *Follow   `json:"follow,omitempty" bson:"follow"`
	Follows []*Follow `json:"follows,omitempty" bson:"follows"`
	Errors  []*Error  `json:"errors,omitempty" bson:"errors"`
}
