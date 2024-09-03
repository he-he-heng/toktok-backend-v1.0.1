package handler

import "toktok-backend-v1.0.1/internal/core/port"

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	userHandler := UserHandler{
		userService: userService,
	}

	return &userHandler
}
