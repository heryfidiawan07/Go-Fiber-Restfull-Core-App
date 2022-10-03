package router

import (
	"api-fiber-gorm/controler"
	"api-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controler.Login)
	auth.Post("/register", controler.Register)
	auth.Post("/refresh-token", controler.RefreshToken)
	auth.Get("/me", middleware.Auth(), controler.Me)

	// User
	user := api.Group("/user")
	user.Get("/", middleware.Auth(), controler.UserIndex)
	user.Get("/:id", controler.UserShow)
	user.Post("/", controler.UserStore)
	user.Put("/:id", controler.UserUpdate)
	user.Delete("/:id", controler.UserDestroy)
	// user.Patch("/:id", middleware.Protected(), controler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), controler.DeleteUser)

	// Role
	role := api.Group("/role")
	role.Get("/", controler.RoleIndex)
	role.Get("/:id", controler.RoleShow)
	role.Post("/", controler.RoleStore)
	role.Put("/:id", controler.RoleUpdate)
	role.Delete("/:id", controler.RoleDestroy)

	// Permission
	permission := api.Group("/permission")
	permission.Get("/", controler.PermissionIndex)
}
