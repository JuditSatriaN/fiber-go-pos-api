package member

import (
	"database/sql"

	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
)

const queryGetAllMember = `
	SELECT id, name, phone
	FROM members
	ORDER BY id
`

func GetAllMember(ctx *fiber.Ctx) ([]model.Member, error) {
	var members []model.Member
	db := postgresPkg.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &members, queryGetAllMember); err != nil {
		return members, err
	}

	return members, nil
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
		if err == sql.ErrNoRows {
			return member, false, nil
		}

		return member, false, err
	}

	return member, true, nil
}

const queryInsertMember = `
	INSERT INTO members (name, phone)
	VALUES (:name, :phone)
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
