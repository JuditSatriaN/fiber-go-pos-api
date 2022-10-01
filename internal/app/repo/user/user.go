package user

import (
	"database/sql"

	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
)

const queryGetAllUser = `
	SELECT user_id, user_name, full_name, password, is_admin
	FROM users
	ORDER BY user_id
`

func GetAllUser(ctx *fiber.Ctx) ([]model.User, error) {
	var users []model.User
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &users, queryGetAllUser); err != nil {
		return users, err
	}

	return users, nil
}

const queryGetUserByUserID = `
	SELECT user_id, user_name, full_name, password, is_admin
	FROM users
	WHERE user_id = $1
`

func GetUserByUserID(ctx *fiber.Ctx, userID string) (model.User, bool, error) {
	var user model.User
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &user, queryGetUserByUserID, userID); err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}

		return user, false, err
	}

	return user, true, nil
}

const queryGetUserByUserName = `
	SELECT user_id, user_name, full_name, password, is_admin
	FROM users
	WHERE user_name = $1
`

func GetUserByUserName(ctx *fiber.Ctx, userName string) (model.User, error) {
	var user model.User
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &user, queryGetUserByUserName, userName); err != nil {
		if err == sql.ErrNoRows {
			return user, constant.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

const queryInsertUser = `
	INSERT INTO users (user_id, user_name, full_name, password, is_admin)
	VALUES (:user_id, :user_name, :full_name, :password, :is_admin)
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
	WHERE user_id = :user_id
`

func UpdateUser(ctx *fiber.Ctx, user model.User) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpdateUser, user)
	return err
}

const queryDeleteUser = `
	DELETE FROM users
	WHERE user_id = $1
`

func DeleteUser(ctx *fiber.Ctx, userID string) error {
	db := postgresPkg.GetPgConn()
	_, err := db.ExecContext(ctx.Context(), queryDeleteUser, userID)
	return err
}
