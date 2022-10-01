package user

import (
	"github.com/gofiber/fiber/v2"

	userHandler "github.com/fiber-go-pos-api/internal/app/handler/user"
)

// BuildUserAPI : API to handle user
func BuildUserAPI(api fiber.Router) {
	api.Get("/users", userHandler.GetAllUserHandler)
	api.Post("/user/insert", userHandler.InsertUserHandler)
	api.Post("/user/update", userHandler.UpdateUserHandler)
	api.Post("/user/delete", userHandler.DeleteUserHandler)
	api.Post("/user/upsert", userHandler.UpsertUserHandler)
}
