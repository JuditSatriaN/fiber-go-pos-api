package user

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	userUC "github.com/fiber-go-pos-api/internal/app/usecase/user"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
)

// GetAllUserHandler : Get List Of User
func GetAllUserHandler(ctx *fiber.Ctx) error {
	users, err := userUC.GetAllUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, users)
}

// GetAllDTUserHandler : Get List Of User For Datatable
func GetAllDTUserHandler(ctx *fiber.Ctx) error {
	users, err := userUC.GetAllUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildDatatableRes(ctx, int64(len(users)), users)
}

// InsertUserHandler : Insert User
func InsertUserHandler(ctx *fiber.Ctx) error {
	var user model.User

	if err := requestPkg.ValidateRequest(ctx, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.InsertUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil disimpan")
}

// UpdateUserHandler : Update User
func UpdateUserHandler(ctx *fiber.Ctx) error {
	var user model.User

	if err := requestPkg.ValidateRequest(ctx, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.UpdateUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil diubah")
}

// DeleteUserHandler : Delete User
func DeleteUserHandler(ctx *fiber.Ctx) error {
	var user model.User

	if err := requestPkg.ValidateRequest(ctx, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.DeleteUser(ctx, user.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data user berhasil dihapus")
}

// UpsertUserHandler : Upsert User
func UpsertUserHandler(ctx *fiber.Ctx) error {
	var user model.User

	if err := requestPkg.ValidateRequest(ctx, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := userUC.UpsertUser(ctx, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, user)
}
