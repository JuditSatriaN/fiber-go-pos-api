package login

import (
	"github.com/gofiber/fiber/v2"

	loginHandler "github.com/fiber-go-pos-api/internal/app/handler/login"
)

// BuildLoginAPI : API to handle login
func BuildLoginAPI(api fiber.Router) {
	api.Post("/login", loginHandler.ProcessLoginHandler)
}
