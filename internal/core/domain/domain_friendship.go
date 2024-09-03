package domain

import "gorm.io/gorm"

type FriendshipStatusType string

const (
	Pending  FriendshipStatusType = "PENDING"
	Accepted FriendshipStatusType = "ACCEPTED"
	Rejected FriendshipStatusType = "REJECTED"
	Removed  FriendshipStatusType = "REMOVED"
)

type Friendship struct {
	gorm.Model

	// user < -- 1 -- 친구추가 -- N --> frindship
	UserID uint `gorm:"not null"`

	// user < -- 1 -- 친구추가 -- N --> frindship
	FriendID uint `gorm:"not null"`

	Status FriendshipStatusType `gorm:"default:PENDING;not null"`

	// friendship <-- 1 -- -- N --> message
	Messages []Message
}
