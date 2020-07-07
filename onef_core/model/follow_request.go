package model

type FollowRequest struct {
	Id           int64 `json:"id,omitempty" bson:"id"`
	CreatorId    int32 `json:"creator_id,omitempty" bson:"creator_id"`
	TargetUserId int32 `json:"target_user_id,omitempty" bson:"target_user_id"`
	Creator      *User `json:"creator,omitempty" bson:"creator"`
	TargetUser   *User `json:"target_user,omitempty" bson:"target_user"`
}

type FollowRequestRequest struct {
}

type FollowRequestResponse struct {
	FollowRequest  *FollowRequest   `json:"followrequest,omitempty" bson:"followrequest"`
	FollowRequests []*FollowRequest `json:"followrequests,omitempty" bson:"followrequests"`
	Errors         []*Error         `json:"errors,omitempty" bson:"errors"`
}
