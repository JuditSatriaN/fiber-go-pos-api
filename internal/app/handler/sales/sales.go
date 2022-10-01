package sales

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	salesUC "github.com/fiber-go-pos-api/internal/app/usecase/sales"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
)

func InsertSalesHandler(ctx *fiber.Ctx) error {
	var sales model.Sales

	if err := requestPkg.ValidateRequest(ctx, &sales); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := salesUC.InsertSales(ctx, sales); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data sales berhasil disimpan")
}
