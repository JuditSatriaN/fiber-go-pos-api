package user

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
	errorPkg "github.com/fiber-go-pos-api/internal/pkg/error"
)

const queryGetAllUser = `
	SELECT id, shop_id, user_name, full_name, password, is_admin
	FROM users
	WHERE shop_id = $1
	AND ($2 = '' OR value_text_search @@ plainto_tsquery($2))
	ORDER BY id
	LIMIT $3
	OFFSET $4
`

func GetAllUser(ctx *fiber.Ctx, shopID int64, search string, limit int, offset int) ([]model.User, error) {
	var users []model.User
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &users, queryGetAllUser, shopID, search, limit, offset); err != nil {
		return users, err
	}

	return users, nil
}

const queryGetTotalDataUsers = `
	SELECT COUNT(*) OVER (ROWS BETWEEN CURRENT ROW AND 1000 FOLLOWING) AS total_count
	FROM users
	WHERE shop_id = $1 
  	AND $2 = '' OR value_text_search @@ plainto_tsquery($2)
	LIMIT 1
`

func GetTotalDataUsers(ctx *fiber.Ctx, shopID int64, search string) (int64, error) {
	var totalData int64
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &totalData, queryGetTotalDataUsers, shopID, search); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return totalData, nil
		}
		return totalData, err
	}

	return totalData, nil
}

const queryGetUserByUserID = `
	SELECT id, shop_id, user_name, full_name, password, is_admin
	FROM users
	WHERE id = $1
`

func GetUserByUserID(ctx *fiber.Ctx, ID int64) (model.User, bool, error) {
	var user model.User
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &user, queryGetUserByUserID, ID); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return user, false, nil
		}

		return user, false, err
	}

	return user, true, nil
}

const queryGetUserByUserName = `
	SELECT id, user_name, full_name, password, is_admin
	FROM users
	WHERE user_name = $1
`

func GetUserByUserName(ctx *fiber.Ctx, userName string) (model.User, error) {
	var user model.User
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &user, queryGetUserByUserName, userName); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return user, constant.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

const queryInsertUser = `
	INSERT INTO users (id, shop_id, user_name, full_name, password, is_admin)
	VALUES (:id, :shop_id, :user_name, :full_name, :password, :is_admin)
`

func InsertUser(ctx *fiber.Ctx, user model.User) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryInsertUser, user)
	return err
}

const queryUpdateUser = `
	UPDATE users SET
		user_name = :user_name,
		full_name = :full_name,
		password = :password,
		is_admin = :is_admin,
		update_time = NOW()
	WHERE id = :id
`

func UpdateUser(ctx *fiber.Ctx, user model.User) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpdateUser, user)
	return err
}

const queryDeleteUser = `
	DELETE FROM users
	WHERE id = $1
`

func DeleteUser(ctx *fiber.Ctx, ID int64) error {
	db := postgresPkg.GetPgConn()
	_, err := db.ExecContext(ctx.Context(), queryDeleteUser, ID)
	return err
}
