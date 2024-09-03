package router

import "toktok-backend-v1.0.1/internal/adapter/presentation/handler"

type HandlerSet struct {
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
}
