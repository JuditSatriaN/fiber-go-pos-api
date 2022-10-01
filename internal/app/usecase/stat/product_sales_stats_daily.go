package stat

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	statRepo "github.com/fiber-go-pos-api/internal/app/repo/stat"
)

func GetTop3ProductSalesDaily(ctx *fiber.Ctx) ([]model.ProductSalesStatsDaily, error) {
	results, err := statRepo.GetProductSalesStatsDaily(ctx, 3)
	if err != nil {
		return []model.ProductSalesStatsDaily{}, err
	}

	return results, nil
}
