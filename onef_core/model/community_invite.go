package model

type CommunityInvite struct {
	Id            int64      `json:"id,omitempty" bson:"id"`
	CreatorId     int64      `json:"creator_id,omitempty" bson:"creator_id"`
	InvitedUserId int64      `json:"invited_user_id,omitempty" bson:"invited_user_id"`
	CommunityId   int64      `json:"community_id,omitempty" bson:"community_id"`
	Creator       *User      `json:"creator,omitempty" bson:"creator"`
	InvitedUser   *User      `json:"invited_user,omitempty" bson:"invited_user"`
	Community     *Community `json:"community,omitempty" bson:"community"`
}

type CommunityInviteRequest struct {
}

type CommunityInviteResponse struct {
	CommunityInvite  *CommunityInvite   `json:"communityinvite,omitempty" bson:"communityinvite"`
	CommunityInvites []*CommunityInvite `json:"communityinvites,omitempty" bson:"communityinvites"`
	Errors           []*Error           `json:"errors,omitempty" bson:"errors"`
}
