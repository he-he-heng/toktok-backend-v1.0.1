package router

import (
	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/adapter/presentation/sniff"
	"toktok-backend-v1.0.1/internal/config"
)

type Router struct {
	handlerSet HandlerSet
	config     *config.Config
	app        *fiber.App
}

func NewRouter(config *config.Config, handlerSet HandlerSet, middlewareSet MiddlewareSet) *Router {

	// Create Configuration  for fiber
	fiberConfig := fiber.Config{
		ErrorHandler: sniff.CustomErrorHandler,
	}

	// Create fiber instance
	app := fiber.New(fiberConfig)

	// set middlware

	// set api endpoint
	v1 := app.Group("/v1")
	{
		api := v1.Group("/api")
		{

			auth := api.Group("/auth")
			{
				auth.Post("/login", handlerSet.AuthHandler.Login)
				auth.Post("/refresh", handlerSet.AuthHandler.Refresh)
				auth.Post("/validation", middlewareSet.GuardMiddleware.TokenValidate, handlerSet.AuthHandler.Validation)
			}

			users := api.Group("/users")
			{
				users.Post("/", handlerSet.UserHandler.CreateUser)
				users.Get("/:id", middlewareSet.GuardMiddleware.FilterValidAccess, handlerSet.UserHandler.GetUsers)
				users.Put("/:id", middlewareSet.GuardMiddleware.FilterValidAccess, handlerSet.UserHandler.UpdateUser)
				users.Delete("/:id", middlewareSet.GuardMiddleware.FilterValidAccess, handlerSet.UserHandler.DeleteUser)
			}
		}
	}

	// Register Router
	router := Router{
		handlerSet: handlerSet,
		config:     config,
		app:        app,
	}

	return &router

}

func (r *Router) Listen() error {
	port := ":8080"
	if r.config.Http.Port == "" {
		port = r.config.Http.Port
	}

	return r.app.Listen(port)
}
