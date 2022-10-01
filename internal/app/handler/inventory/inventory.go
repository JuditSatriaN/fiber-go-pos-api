package inventory

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	inventoryUC "github.com/fiber-go-pos-api/internal/app/usecase/inventory"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
)

// GetInventoryHandler : Get List Data Of Inventory
func GetInventoryHandler(ctx *fiber.Ctx) error {
	page, limit, err := requestPkg.BuildPageAndLimit(ctx)
	if err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	search := requestPkg.BuildSearchRequest(ctx)
	result, err := inventoryUC.GetAllInventory(ctx, page, limit, search)
	if err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusOK,
		Data:         result.Data,
		Metadata:     responsePkg.BuildMetaDataResponse(page, limit, int(result.Total), nil),
	})
}

// SearchInventoryHandler : Search all inventory by param
func SearchInventoryHandler(ctx *fiber.Ctx) error {
	search := requestPkg.BuildSearchRequest(ctx)

	result, err := inventoryUC.SearchInventoryByParam(ctx, search)
	if err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{Data: result})
}

// InsertInventoryHandler : Insert inventory
func InsertInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := requestPkg.ValidateRequest(ctx, &inventory); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	if err := inventoryUC.InsertInventory(ctx, inventory); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusCreated,
		Message:      "Data inventory berhasil ditambahkan",
	})
}

// UpdateInventoryHandler : Update inventory
func UpdateInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := requestPkg.ValidateRequest(ctx, &inventory); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	if err := inventoryUC.UpdateInventory(ctx, inventory); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusOK,
		Message:      "Data inventory berhasil diubah",
	})
}

func DeleteInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := requestPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.DeleteInventory(ctx, inventory.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data inventory berhasil dihapus")
}

func UpsertInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := requestPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.UpsertInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{Data: inventory})
}

func UpdateStockInventoryHandler(ctx *fiber.Ctx) error {
	var inventory model.Inventory

	if err := requestPkg.ValidateRequest(ctx, &inventory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := inventoryUC.UpdateStockInventory(ctx, inventory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data update stock berhasil diubah")
}
