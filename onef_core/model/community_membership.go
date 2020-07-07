package model

type CommunityMembership struct {
	Id              int64      `json:"id,omitempty" bson:"id"`
	CreatorId       int64      `json:"creator_id,omitempty" bson:"creator_id"`
	CommunityId     int64      `json:"community_id,omitempty" bson:"community_id"`
	IsAdministrator bool       `json:"is_administrator,omitempty" bson:"is_administrator"`
	IsModerator     bool       `json:"is_moderator,omitempty" bson:"is_moderator"`
	Creator         *User      `json:"creator,omitempty" bson:"creator"`
	Community       *Community `json:"community,omitempty" bson:"community"`
}

type CommunityMembershipRequest struct {
}

type CommunityMembershipResponse struct {
	CommunityMembership  *CommunityMembership   `json:"communitymembership,omitempty" bson:"communitymembership"`
	CommunityMemberships []*CommunityMembership `json:"communitymemberships,omitempty" bson:"communitymemberships"`
	Errors               []*Error               `json:"errors,omitempty" bson:"errors"`
}
