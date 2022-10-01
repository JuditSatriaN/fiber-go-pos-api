package router

import (
	"github.com/gofiber/fiber/v2"

	inventoryRoute "github.com/fiber-go-pos-api/internal/app/router/inventory"
	loginRoute "github.com/fiber-go-pos-api/internal/app/router/login"
	memberRoute "github.com/fiber-go-pos-api/internal/app/router/member"
	productRoute "github.com/fiber-go-pos-api/internal/app/router/product"
	salesRoute "github.com/fiber-go-pos-api/internal/app/router/sales"
	statRoute "github.com/fiber-go-pos-api/internal/app/router/stat"
	unitRoute "github.com/fiber-go-pos-api/internal/app/router/unit"
	userRoute "github.com/fiber-go-pos-api/internal/app/router/user"
)

// BuildAPIRouter : Function to handle all API in this project
func BuildAPIRouter(app *fiber.App) {
	apiGroup := app.Group("/api")
	userRoute.BuildUserAPI(apiGroup)
	unitRoute.BuildUnitAPI(apiGroup)
	loginRoute.BuildLoginAPI(apiGroup)
	salesRoute.BuildSalesAPI(apiGroup)
	memberRoute.BuildMemberAPI(apiGroup)
	productRoute.BuildProductAPI(apiGroup)
	inventoryRoute.BuildInventoryAPI(apiGroup)
	statRoute.BuildProductSalesStatsDailyAPI(apiGroup)
}
