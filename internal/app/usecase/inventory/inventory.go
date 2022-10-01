package inventory

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	inventoryRepo "github.com/fiber-go-pos-api/internal/app/repo/inventory"
	productRepo "github.com/fiber-go-pos-api/internal/app/repo/product"
	statRepo "github.com/fiber-go-pos-api/internal/app/repo/stat"
	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
)

// GetAllInventory : Get List Of Data Inventory
func GetAllInventory(ctx *fiber.Ctx, page int, limit int, search string) (model.ListInventoryDataResponse, error) {
	offset := requestPkg.BuildOffset(page, limit)

	inventories, err := inventoryRepo.GetALlInventory(ctx, search, limit, offset)
	if err != nil {
		return model.ListInventoryDataResponse{}, err
	}

	totalData, err := inventoryRepo.GetTotalDataInventory(ctx, search)
	if err != nil {
		return model.ListInventoryDataResponse{}, err
	}

	return model.ListInventoryDataResponse{
		Total: totalData,
		Data:  inventories,
	}, nil
}

// SearchInventoryByParam : Search Inventory By Parameter Request
func SearchInventoryByParam(ctx *fiber.Ctx, search string) ([]model.Inventory, error) {
	product, found, err := productRepo.GetProductByPLUOrBarcode(ctx, search)
	if err != nil {
		return []model.Inventory{}, err
	}

	if !found {
		return []model.Inventory{}, constant.ErrNoDataFound
	}

	results, err := inventoryRepo.GetInventoryByPLU(ctx, product.ID)
	if err != nil {
		return []model.Inventory{}, err
	}

	return results, nil
}

// GetInventoryByID : Get Inventory By ID
func GetInventoryByID(ctx *fiber.Ctx, ID int64) (model.Inventory, error) {
	inventory, found, err := inventoryRepo.GetInventoryByID(ctx, ID)
	if err != nil {
		return model.Inventory{}, err
	}

	if !found {
		return model.Inventory{}, constant.ErrNoDataFound
	}

	return inventory, nil
}

// InsertInventory : Insert Inventory
func InsertInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := inventoryRepo.InsertInventory(tx, inventory); err != nil {
		return err
	}

	err = statRepo.UpdateTotalInventory(tx, model.StoreStats{
		StoreID:        constant.DefaultStoreID,
		TotalInventory: int64(constant.Increment),
	})

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	if _, err := GetInventoryByID(ctx, inventory.ID); err != nil {
		return err
	}
	return inventoryRepo.UpdateInventory(ctx, inventory)
}

func DeleteInventory(ctx *fiber.Ctx, ID int64) error {
	if _, err := GetInventoryByID(ctx, ID); err != nil {
		return err
	}

	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := inventoryRepo.DeleteInventory(tx, ID); err != nil {
		return err
	}

	err = statRepo.UpdateTotalInventory(tx, model.StoreStats{
		StoreID:        constant.DefaultStoreID,
		TotalInventory: int64(constant.Decrement),
	})
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpsertInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	_, found, err := inventoryRepo.GetInventoryByID(ctx, inventory.ID)
	if err != nil {
		return err
	}

	if !found || inventory.ID == 0 {
		if err := InsertInventory(ctx, inventory); err != nil {
			return err
		}
	} else {
		if err := UpdateInventory(ctx, inventory); err != nil {
			return err
		}
	}

	return nil
}

func UpdateStockInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	if _, err := GetInventoryByID(ctx, inventory.ID); err != nil {
		return err
	}

	return inventoryRepo.UpdateStockInventory(ctx, inventory)
}
