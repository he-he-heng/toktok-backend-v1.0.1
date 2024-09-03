package domain

import "gorm.io/gorm"

type UserRoleType string

const (
	Guest   UserRoleType = "GUEST"
	Manager UserRoleType = "MANAGER"
	Admin   UserRoleType = "ADMIN"
)

type User struct {
	gorm.Model

	LoginID  string
	Password string       `gorm:"type:varchar(255);not null"`
	Email    *string      `gorm:"type:varchar(255)"`
	Role     UserRoleType `gorm:"default:GUEST;not null"`

	// user < -- 1 -- 친구추가 -- N --> frindship
	UserFriendships []Friendship `gorm:"foreignKey:UserID"`

	// user < -- 1 -- 친구추가 -- N --> frindship
	FriendFriendships []Friendship `gorm:"foreignKey:FriendID"`

	// user < -- 1 -- -- N -- > message
	Messages []Message

	// user < -- 1 -- -- 1 -- > avatar
	Avatar Avatar
}
