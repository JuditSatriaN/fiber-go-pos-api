package unit

import (
	"fmt"

	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	unitRepo "github.com/fiber-go-pos-api/internal/app/repo/unit"
)

func GetAllUnit(ctx *fiber.Ctx) ([]model.Unit, error) {
	unit, err := unitRepo.GetAllUnit(ctx)
	if err != nil {
		return []model.Unit{}, err
	}

	return unit, nil
}

func GetUnitByID(ctx *fiber.Ctx, ID int64) (model.Unit, error) {
	unit, found, err := unitRepo.GetUnitByID(ctx, ID)
	if err != nil {
		return model.Unit{}, err
	}

	if !found {
		return model.Unit{}, fmt.Errorf("unit dengan id : %d tidak ditemukan", ID)
	}

	return unit, nil
}

func InsertUnit(ctx *fiber.Ctx, unit model.Unit) error {
	return unitRepo.InsertUnit(ctx, unit)
}

func UpdateUnit(ctx *fiber.Ctx, unit model.Unit) error {
	if _, err := GetUnitByID(ctx, unit.ID); err != nil {
		return err
	}

	return unitRepo.UpdateUnit(ctx, unit)
}

func DeleteUnit(ctx *fiber.Ctx, ID int64) error {
	if _, err := GetUnitByID(ctx, ID); err != nil {
		return err
	}

	return unitRepo.DeleteUnit(ctx, ID)
}

func UpsertUnit(ctx *fiber.Ctx, unit model.Unit) error {
	_, found, err := unitRepo.GetUnitByID(ctx, unit.ID)
	if err != nil {
		return err
	}

	if !found || unit.ID == 0 {
		return unitRepo.InsertUnit(ctx, unit)
	} else {
		return unitRepo.UpdateUnit(ctx, unit)
	}
}
