package user

import (
	"fmt"

	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	userRepo "github.com/fiber-go-pos-api/internal/app/repo/user"
	passwordPkg "github.com/fiber-go-pos-api/internal/pkg/password"
)

func GetAllUser(ctx *fiber.Ctx) ([]model.User, error) {
	users, err := userRepo.GetAllUser(ctx)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func GetUserByUserID(ctx *fiber.Ctx, ID int64) (model.User, error) {
	user, found, err := userRepo.GetUserByUserID(ctx, ID)
	if err != nil {
		return model.User{}, err
	}

	if !found {
		return model.User{}, fmt.Errorf("user dengan nama : %s tidak ditemukan", user.UserName)
	}

	return user, nil
}

func InsertUser(ctx *fiber.Ctx, user model.User) error {
	user.Password, _ = passwordPkg.HashPassword(user.Password)
	return userRepo.InsertUser(ctx, user)
}

func UpdateUser(ctx *fiber.Ctx, user model.User) error {
	userDB, err := GetUserByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	// replace to existing data
	if userDB.Password != user.Password {
		user.Password, _ = passwordPkg.HashPassword(user.Password)
	}

	return userRepo.UpdateUser(ctx, user)
}

func DeleteUser(ctx *fiber.Ctx, ID int64) error {
	if _, err := GetUserByUserID(ctx, ID); err != nil {
		return err
	}

	return userRepo.DeleteUser(ctx, ID)
}

func UpsertUser(ctx *fiber.Ctx, user model.User) error {
	userDB, found, err := userRepo.GetUserByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	// replace to existing data
	if userDB.Password != user.Password {
		user.Password, _ = passwordPkg.HashPassword(user.Password)
	}

	if !found {
		return userRepo.InsertUser(ctx, user)
	} else {
		return userRepo.UpdateUser(ctx, user)
	}
}
