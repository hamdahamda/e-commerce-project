package router

import (
	"ecommerce-app/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	// routes activity
	app.Get("/activities", handler.GetActivities)
	app.Post("/activities", handler.CreateActivity)
	app.Put("/activities/:id", handler.UpdateActivity)
	app.Delete("/activities/:id", handler.DeleteActivity)

	// routes user
	app.Get("/user", handler.GetUser)
	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)

	// routes user_roles
	app.Get("/user_role", handler.GetUsersRole)
	app.Post("/user_role", handler.CreateUserRole)
	app.Put("/user_role/:id", handler.UpdateUserRole)
	app.Delete("/user_role/:id", handler.DeleteUserRole)

	// routes user_roles_map
	app.Get("/user_role_map", handler.GetUsersRoleMap)
	app.Post("/user_role_map", handler.CreateUserRoleMap)
	app.Put("/user_role_map/:id", handler.UpdateUserRoleMap)
	app.Delete("/user_role_map/:id", handler.DeleteUserRoleMap)

	// routes categories
	app.Get("/categories", handler.GetCategories)
	app.Post("/categories", handler.CreateCategory)
	app.Put("/categories/:id", handler.UpdateCategory)
	app.Delete("/categories/:id", handler.DeleteCategory)

	return app
}
