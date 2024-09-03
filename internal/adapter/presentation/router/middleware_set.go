package router

import "toktok-backend-v1.0.1/internal/adapter/presentation/middleware"

type MiddlewareSet struct {
	GuardMiddleware *middleware.GuardMiddleware
}
