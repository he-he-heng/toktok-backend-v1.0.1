package domain

type TokenType string

const (
	AccessToken  TokenType = "access-token"
	RefreshToken TokenType = "refresh-token"
)

type TokenPayload struct {

	// 토큰을 발급한 사람의 userID
	Iss int

	// 토큰 만료 시간 (unix timestamp로 표현)
	Exp int64

	// 토큰이 발급된 시간 (unix timestamp로 표현)
	Ita int64

	TokenType TokenType

	Role UserRoleType
}
