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
	var member model.MemberInsertPayload

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	err := memberUC.InsertMember(ctx, model.Member{
		ShopID:  member.ShopID,
		Name:    member.Name,
		Phone:   member.Phone,
		Address: member.Address,
	})
	if err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusCreated,
		Message:      "Data member berhasil ditambahkan",
	})
}

func UpdateMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	if err := memberUC.UpdateMember(ctx, member); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusCreated,
		Message:      "Data member berhasil diubah",
	})
}

func DeleteMemberHandler(ctx *fiber.Ctx) error {
	var member model.MemberDeletePayload

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	if err := memberUC.DeleteMember(ctx, member.ID); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusCreated,
		Message:      "Data member berhasil dihapus",
	})
}

func UpsertMemberHandler(ctx *fiber.Ctx) error {
	var member model.Member

	if err := requestPkg.ValidateRequest(ctx, &member); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusBadRequest,
			Message:      err.Error(),
		})
	}

	if err := memberUC.UpsertMember(ctx, member); err != nil {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusInternalServerError,
			Message:      err.Error(),
		})
	}

	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusOK,
		Message:      "Data member berhasil diubah",
	})
}
