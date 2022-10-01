package product

import (
	"github.com/gofiber/fiber/v2"

	productHandler "github.com/fiber-go-pos-api/internal/app/handler/product"
)

// BuildProductAPI : API to handle product
func BuildProductAPI(api fiber.Router) {
	api.Get("/dt_products", productHandler.GetALLDTProductHandler)
	api.Get("/products", productHandler.GetALLProductHandler)
	api.Get("/product", productHandler.GetProductByPLUOrBarcode)
	api.Post("/product/insert", productHandler.InsertProductHandler)
	api.Post("/product/update", productHandler.UpdateProductHandler)
	api.Post("/product/delete", productHandler.DeleteProductHandler)
	api.Post("/product/upsert", productHandler.UpsertProductHandler)
}
