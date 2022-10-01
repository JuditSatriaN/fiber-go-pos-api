package inventory

import (
	"github.com/gofiber/fiber/v2"

	inventoryHandler "github.com/fiber-go-pos-api/internal/app/handler/inventory"
)

// BuildInventoryAPI : API to handle inventory product
func BuildInventoryAPI(api fiber.Router) {
	api.Get("/inventory", inventoryHandler.GetInventoryHandler)
	api.Get("/search_inventory", inventoryHandler.SearchInventoryHandler)
	api.Post("/inventory/insert", inventoryHandler.InsertInventoryHandler)
	api.Post("/inventory/update", inventoryHandler.UpdateInventoryHandler)
	api.Post("/inventory/delete", inventoryHandler.DeleteInventoryHandler)
	api.Post("/inventory/upsert", inventoryHandler.UpsertInventoryHandler)
	api.Post("/inventory/update_stock", inventoryHandler.UpdateStockInventoryHandler)
}
