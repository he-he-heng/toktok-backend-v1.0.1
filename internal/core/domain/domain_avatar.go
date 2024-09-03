package domain

type AvatarSexType string

const (
	Male   AvatarSexType = "MALE"
	Female AvatarSexType = "FEMALE"
)

type Avatar struct {
	Sex       AvatarSexType
	Birthday  string
	Mbti      *string
	Picture   string
	Nickanme  string
	Introduce *string

	// user < -- 1 -- -- 1 -- > avatar
	UserID uint
}
