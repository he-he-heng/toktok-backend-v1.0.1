package domain

import "time"

type Message struct {
	// user < -- 1 -- -- N -- > message
	UserID uint

	// friendship <-- 1 -- -- N --> message
	FriendshipID uint

	Content  string
	SentAt   time.Time
	ReadedAt *time.Time
}
