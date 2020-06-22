package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id          int          `json:"id,omitempty" bson:"id"`
	UserName    string       `json:"user_name,omitempty" bsn:"user_name"`
	Email       string       `json:"email,omitempty" bson:"email"`
	Password    string       `json:"password,omitempty" bson:"password"`
	UserProfile *UserProfile `json"profile,omitempty"`
}

type UserProfile struct {
	Id       int    `json:"id,omitempty" bson:"id"`
	Name     string `json:"name,omitempty" bsn:"name"`
	Avatar   string `json:"avatar,omitempty" bson:"avatar"`
	Cover    string `json:"cover,omitempty" bson:"cover"`
	Bio      string `json:"bio,omitempty" bson:"bio"`
	Url      string `json:"url,omitempty" bson:"url"`
	Location string `json:"location,omitempty" bson:"location"`
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
	User   *User    `json"user,omitempty"`
	Users  []*User  `json:"users,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
}

type Token struct {
	Token  string   `jon:"token,omitempty"`
	Valid  bool     `json:"valid,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
}
