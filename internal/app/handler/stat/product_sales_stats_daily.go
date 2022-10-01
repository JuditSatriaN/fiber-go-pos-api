package stat

import (
	"github.com/gofiber/fiber/v2"

	statUC "github.com/fiber-go-pos-api/internal/app/usecase/stat"
	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
)

func GetTop3ProductSalesDailyHandler(ctx *fiber.Ctx) error {
	results, err := statUC.GetTop3ProductSalesDaily(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, results)
}
