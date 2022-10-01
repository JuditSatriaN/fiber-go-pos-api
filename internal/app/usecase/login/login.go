package login

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	userRepo "github.com/fiber-go-pos-api/internal/app/repo/user"
	jwtPkg "github.com/fiber-go-pos-api/internal/pkg/jwt"
	passwordPkg "github.com/fiber-go-pos-api/internal/pkg/password"
)

func ProcessLoginForm(ctx *fiber.Ctx, req model.LoginRequest) (model.LoginResponse, error) {
	// Initialization variable
	var res model.LoginResponse

	data, err := userRepo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return res, constant.ErrInvalidLogin
	}

	// Check hash password
	if !passwordPkg.CheckPasswordHash(req.Password, data.Password) {
		return res, constant.ErrInvalidLogin
	}

	// Create login token to set in cookie
	token, err := jwtPkg.CreateJWTToken(constant.JWTRequest{
		UserID:  data.UserID,
		Name:    data.UserName,
		IsAdmin: data.IsAdmin,
	})
	if err != nil {
		return res, constant.ErrInvalidLogin
	}

	return model.LoginResponse{
		UserID:          data.UserID,
		UserName:        data.UserName,
		FullName:        data.FullName,
		IsAdmin:         data.IsAdmin,
		JWTAccessToken:  token.AccessToken,
		JWTRefreshToken: token.RefreshToken,
	}, nil
}
