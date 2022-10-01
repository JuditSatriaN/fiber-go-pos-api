package unit

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	unitUC "github.com/fiber-go-pos-api/internal/app/usecase/unit"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
)

func GetAllUnitHandler(ctx *fiber.Ctx) error {
	units, err := unitUC.GetAllUnit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, units)
}

func GetAllDTUnitHandler(ctx *fiber.Ctx) error {
	units, err := unitUC.GetAllUnit(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildDatatableRes(ctx, int64(len(units)), units)
}

func InsertUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := requestPkg.ValidateRequest(ctx, &unit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := unitUC.InsertUnit(ctx, unit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data unit berhasil disimpan")
}

func UpdateUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := requestPkg.ValidateRequest(ctx, &unit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := unitUC.UpdateUnit(ctx, unit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data unit berhasil diubah")
}

func DeleteUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := requestPkg.ValidateRequest(ctx, &unit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := unitUC.DeleteUnit(ctx, unit.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data unit berhasil dihapus")
}

func UpsertUnitHandler(ctx *fiber.Ctx) error {
	var unit model.Unit

	if err := requestPkg.ValidateRequest(ctx, &unit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := unitUC.UpsertUnit(ctx, unit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, unit)
}
