package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id       int          `json:"id,omitempty" bson:"id"`
	UserName string       `json:"username,omitempty" bson:"username"`
	Email    string       `json:"email,omitempty" bson:"email"`
	Password string       `json:"password,omitempty" bson:"password"`
	Profile  *UserProfile `json:"profile,omitempty" bson:"profile"`
}

type UserProfile struct {
	Id                    int    `json:"id,omitempty" bson:"id"`
	Name                  string `json:"name,omitempty" bson:"name"`
	Avatar                string `json:"avatar,omitempty" bson:"avatar"`
	Cover                 string `json:"cover,omitempty" bson:"cover"`
	Bio                   string `json:"bio,omitempty" bson:"bio"`
	Url                   string `json:"url,omitempty" bson:"url"`
	Location              string `json:"location,omitempty" bson:"location"`
	IsOfLegalAge          string `json:"is_of_legal_age,omitempty" bson:"is_of_legal_age"`
	FollowersCountVisible string `json:"followers_count_visible,omitempty" bson:"followers_count_visible"`
	CommunityPostsVisible string `json:"community_posts_visible,omitempty" bson:"community_posts_visible"`
}

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

type UserRequest struct {
}

type UserResponse struct {
	User   *User    `json:"user,omitempty" bson:"user"`
	Users  []*User  `json:"users,omitempty" bson:"users"`
	Errors []*Error `json:"errors,omitempty" bson:"errors"`
}

type Token struct {
	UserName string   `json:"username,omitempty" bson:"username"`
	Email    string   `json:"email,omitempty" bson:"email"`
	Token    string   `json:"token,omitempty" bson:"token"`
	Valid    bool     `json:"valid,omitempty" bson:"valid"`
	Errors   []*Error `json:"errors,omitempty" bson:"errors"`
}
