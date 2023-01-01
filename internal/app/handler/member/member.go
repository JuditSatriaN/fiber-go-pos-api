package member

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	memberUC "github.com/fiber-go-pos-api/internal/app/usecase/member"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
)

func GetAllMemberHandler(ctx *fiber.Ctx) error {
	shopID, err := requestPkg.BuildShopIDRequest(ctx)
	if err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	page, limit, search, err := requestPkg.BuildPageLimitAndSearch(ctx)
	if err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	result, err := memberUC.GetAllMember(ctx, shopID, page, limit, search)
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

func InsertMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.InsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil disimpan")
}

func UpdateMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.UpdateMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Data member berhasil diubah")
}

func DeleteMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.DeleteMember(ctx, member.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendString("Data Member berhasil dihapus")
}

func UpsertMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := memberUC.UpsertMember(ctx, member); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, member)
}
