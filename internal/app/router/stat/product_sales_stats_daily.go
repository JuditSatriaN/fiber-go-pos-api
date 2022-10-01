package stat

import (
	"github.com/gofiber/fiber/v2"

	statHandler "github.com/fiber-go-pos-api/internal/app/handler/stat"
)

// BuildProductSalesStatsDailyAPI : API to handle product_sales_stats_daily
func BuildProductSalesStatsDailyAPI(api fiber.Router) {
	api.Get("/product_sales_stats_daily", statHandler.GetTop3ProductSalesDailyHandler)
}
