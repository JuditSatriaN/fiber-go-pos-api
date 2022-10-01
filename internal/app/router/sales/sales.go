package sales

import (
	"github.com/gofiber/fiber/v2"

	salesHandler "github.com/fiber-go-pos-api/internal/app/handler/sales"
)

// BuildSalesAPI : API to handle sales
func BuildSalesAPI(api fiber.Router) {
	api.Post("/sales", salesHandler.InsertSalesHandler)
}
