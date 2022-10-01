package member

import (
	"fmt"

	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	memberRepo "github.com/fiber-go-pos-api/internal/app/repo/member"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
)

func GetAllMember(ctx *fiber.Ctx, shopID int64, page int, limit int, search string) (model.ListMemberDataResponse, error) {
	offset := requestPkg.BuildOffset(page, limit)

	members, err := memberRepo.GetAllMember(ctx, shopID, search, limit, offset)
	if err != nil {
		return model.ListMemberDataResponse{}, err
	}

	totalData, err := memberRepo.GetTotalDataMember(ctx, shopID, search)
	if err != nil {
		return model.ListMemberDataResponse{}, err
	}

	return model.ListMemberDataResponse{
		Total: totalData,
		Data:  members,
	}, nil
}

func GetMemberByID(ctx *fiber.Ctx, ID int64) (model.Member, error) {
	member, found, err := memberRepo.GetMemberByID(ctx, ID)
	if err != nil {
		return model.Member{}, err
	}

	if !found {
		return model.Member{}, fmt.Errorf("member dengan id : %d tidak ditemukan", ID)
	}

	return member, nil
}

func InsertMember(ctx *fiber.Ctx, member model.Member) error {
	return memberRepo.InsertMember(ctx, member)
}

func UpdateMember(ctx *fiber.Ctx, member model.Member) error {
	if _, err := GetMemberByID(ctx, member.ID); err != nil {
		return err
	}

	return memberRepo.UpdateMember(ctx, member)
}

func DeleteMember(ctx *fiber.Ctx, ID int64) error {
	if _, err := GetMemberByID(ctx, ID); err != nil {
		return err
	}

	return memberRepo.DeleteMember(ctx, ID)
}

func UpsertMember(ctx *fiber.Ctx, member model.Member) error {
	_, found, err := memberRepo.GetMemberByID(ctx, member.ID)
	if err != nil {
		return err
	}

	if !found || member.ID == 0 {
		return memberRepo.InsertMember(ctx, member)
	} else {
		return memberRepo.UpdateMember(ctx, member)
	}
}
