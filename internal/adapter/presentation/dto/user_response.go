package dto

import "toktok-backend-v1.0.1/internal/core/domain"

// POST v1/api/users

// GET /v1/api/users/{id}
type getUserResponse struct {
	ID        int     `json:"id"`
	LoginID   string  `json:"login_id"`
	Email     *string `json:"email,omitempty"`
	Role      string  `json:"role"`
	CreatedAt string
	UpdatedAt string
}

func GetUserResponseOf(user *domain.User) getUserResponse {
	return getUserResponse{
		ID:        int(user.ID),
		LoginID:   user.LoginID,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt.Format(createdAtFormat),
		UpdatedAt: user.UpdatedAt.Format(updatedAtFormat),
	}
}
