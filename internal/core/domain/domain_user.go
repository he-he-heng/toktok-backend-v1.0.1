package domain

import "gorm.io/gorm"

type UserRoleType string

const (
	Guest   UserRoleType = "GUEST"
	Manager UserRoleType = "MANAGER"
	Admin   UserRoleType = "ADMIN"
)

type UserSexType string

const (
	Male   UserSexType = "MALE"
	Female UserSexType = "FEMALE"
)

type User struct {
	gorm.Model

	LoginID  string
	Password string       `gorm:"type:varchar(255);not null"`
	Email    *string      `gorm:"type:varchar(255)"`
	Role     UserRoleType `gorm:"type:user_role_type;default:GUEST;not null"`

	Sex       UserSexType
	Birthday  string
	Mbti      *string
	Picture   string
	Nickanme  string
	Introduce *string

	// user < -- 1 -- 친구추가 -- N --> frindship
	UserFriendships []Friendship `gorm:"foreignKey:UserID"`

	// user < -- 1 -- 친구추가 -- N --> frindship
	FriendFriendships []Friendship `gorm:"foreignKey:FriendID"`

	// user < -- 1 -- -- N -- > message
	Messages []Message
}
