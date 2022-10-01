package member

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
	errorPkg "github.com/fiber-go-pos-api/internal/pkg/error"
)

const queryGetAllMember = `
	SELECT id, shop_id, name, phone, address
	FROM members
	WHERE shop_id = $1 
	AND ($2 = '' OR value_text_search @@ plainto_tsquery($2))
	ORDER BY id
	LIMIT $3
	OFFSET $4
`

func GetAllMember(ctx *fiber.Ctx, shopID int64, search string, limit int, offset int) ([]model.Member, error) {
	var members []model.Member
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &members, queryGetAllMember, shopID, search, limit, offset); err != nil {
		return members, err
	}

	return members, nil
}

const queryGetTotalDataMember = `
	SELECT COUNT(*) OVER (ROWS BETWEEN CURRENT ROW AND 1000 FOLLOWING) AS total_count
	FROM members
	WHERE shop_id = $1 
	AND $2 = '' OR value_text_search @@ plainto_tsquery($2)
	LIMIT 1
`

func GetTotalDataMember(ctx *fiber.Ctx, shopID int64, search string) (int64, error) {
	var totalData int64
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &totalData, queryGetTotalDataMember, shopID, search); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return totalData, nil
		}
		return totalData, err
	}

	return totalData, nil
}

const queryGetMemberByID = `
	SELECT id, name, phone
	FROM members
	WHERE id = $1
`

func GetMemberByID(ctx *fiber.Ctx, ID int64) (model.Member, bool, error) {
	var member model.Member

	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &member, queryGetMemberByID, ID); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return member, false, nil
		}

		return member, false, err
	}

	return member, true, nil
}

const queryInsertMember = `
	INSERT INTO members (shop_id, name, phone, address)
	VALUES (:shop_id, :name, :phone, :address)
`

func InsertMember(ctx *fiber.Ctx, member model.Member) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryInsertMember, member)
	return err
}

const queryUpdateMember = `
	UPDATE members 
	SET name = :name,
	    phone = :phone,
	    address = :address,
		update_time = NOW()
	WHERE id = :id
`

func UpdateMember(ctx *fiber.Ctx, member model.Member) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpdateMember, member)
	return err
}

const queryDeleteMember = `
	DELETE FROM members
	WHERE id = $1
`

func DeleteMember(ctx *fiber.Ctx, ID int64) error {
	db := postgresPkg.GetPgConn()
	_, err := db.ExecContext(ctx.Context(), queryDeleteMember, ID)
	return err
}
