package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Model
	UserName              string       `json:"username,omitempty" gorm:"column:username" bson:"username"`
	Email                 string       `json:"email,omitempty" gorm:"column:email" bson:"email"`
	Password              string       `json:"password,omitempty" gorm:"column:password" bson:"password"`
	LastLogin             time.Time    `json:"last_login,omitempty" gorm:"column:last_login" bson:"last_login"`
	IsSuperuser           bool         `json:"is_superuser,omitempty" gorm:"column:is_superuser" bson:"is_superuser"`
	IsStaff               bool         `json:"is_staff,omitempty" gorm:"column:is_staff" bson:"is_staff"`
	IsActive              bool         `json:"is_active,omitempty" gorm:"column:is_active" bson:"is_active"`
	DateJoined            time.Time    `json:"date_joined,omitempty" gorm:"column:date_joined" bson:"date_joined"`
	IsEmailVerified       bool         `json:"is_email_verified,omitempty" gorm:"column:is_email_verified" bson:"is_email_verified"`
	AreGuidelinesAccepted bool         `json:"are_guidelines_accepted,omitempty" gorm:"column:are_guidelines_accepted" bson:"are_guidelines_accepted"`
	IsDeleted             bool         `json:"is_deleted,omitempty" gorm:"column:is_deleted" bson:"is_deleted"`
	Visibility            string       `json:"visibility,omitempty" gorm:"column:visibility" bson:"visibility"`
	InviteCount           int          `json:"invite_count,omitempty" gorm:"column:invite_count" bson:"invite_count"`
	LanguageId            *int64       `json:"language_id,omitempty" gorm:"column:language_id" bson:"language_id"`
	TranslationLanguageId *int64       `json:"translation_language_id,omitempty" gorm:"column:translation_language_id" bson:"translation_language_id"`
	Language              *Language    `json:"language,omitempty" gorm:"foreignkey:LanguageId" bson:"language"`
	TranslationLanguage   *Language    `json:"translation_language,omitempty" gorm:"foreignkey:TranslationLanguageId" bson:"translation_language"`
	Profile               *UserProfile `json:"profile,omitempty"`
	ConnectionsCircleId   int64        `json:"connections_circle_id,omitempty" gorm:"_" sql:"-" bson:"connections_circle_id"`
}

type UserProfile struct {
	Model
	UserId                int64  `json:"user_id,omitempty" gorm:"column:user_id" bson:"user_id"`
	Name                  string `json:"name,omitempty" gorm:"column:name" bson:"name"`
	Avatar                string `json:"avatar,omitempty" gorm:"column:avatar" bson:"avatar"`
	Cover                 string `json:"cover,omitempty" gorm:"column:cover" bson:"cover"`
	Bio                   string `json:"bio,omitempty" gorm:"column:bio" bson:"bio"`
	Url                   string `json:"url,omitempty" gorm:"column:url" bson:"url"`
	Location              string `json:"location,omitempty" gorm:"column:location" bson:"location"`
	IsOfLegalAge          bool   `json:"is_of_legal_age,omitempty" gorm:"column:is_of_legal_age" bson:"is_of_legal_age"`
	FollowersCountVisible bool   `json:"followers_count_visible,omitempty" gorm:"column:followers_count_visible" bson:"followers_count_visible"`
	CommunityPostsVisible bool   `json:"community_posts_visible,omitempty" gorm:"column:community_posts_visible" bson:"community_posts_visible"`
	User                  *User  `json:"user,omitempty" gorm:"foreignkey:UserId" bson:"user"`
}

func (User) TableName() string {
	return "user"
}

func (UserProfile) TableName() string {
	return "userprofile"
}

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

type RegisterRequest struct {
	UserName              string `json:"username,omitempty" bson:"username"`
	Email                 string `json:"email,omitempty" bson:"email"`
	Password              string `json:"password,omitempty" bson:"password"`
	Name                  string `json:"name,omitempty" bson:"name"`
	Avatar                string `json:"avatar,omitempty" bson:"avatar"`
	IsOfLegalAge          bool   `json:"is_of_legal_age,omitempty" bson:"is_of_legal_age"`
	AreGuidelinesAccepted bool   `json:"are_guidelines_accepted,omitempty" bson:"are_guidelines_accepted"`
}

type LoginRequest struct {
	UserName string `json:"username,omitempty" bson:"username"`
	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
}

type UserRequest struct {
	UserName string `json:"username,omitempty" bson:"username"`
	Email    string `json:"email,omitempty" bson:"email"`
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

type AuthenticatedUser struct {
	User
	Profile *UserProfile `json:"profile,omitempty"`
}
