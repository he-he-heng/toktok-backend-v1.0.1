package dto

import "toktok-backend-v1.0.1/internal/core/domain"

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginResponseOf(access, refresh string) loginResponse {
	return loginResponse{AccessToken: access, RefreshToken: refresh}
}

type refreshResponse struct {
	AccessToken string `json:"access_token"`
}

func RefreshResponseOf(access string) refreshResponse {
	return refreshResponse{AccessToken: access}
}

type validationResponse struct {
	Iss       int    `json:"iss"`
	Exp       int64  `json:"exp"`
	Ita       int64  `json:"ita"`
	TokenType string `json:"token_type"`
	Role      string `json:"role"`
}

func ValidateResponseOf(payload *domain.TokenPayload) validationResponse {
	return validationResponse{
		Iss:       payload.Iss,
		Exp:       payload.Exp,
		Ita:       payload.Ita,
		TokenType: string(payload.TokenType),
		Role:      string(payload.Role),
	}
}
