package login

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	loginUC "github.com/fiber-go-pos-api/internal/app/usecase/login"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
)

func ProcessLoginHandler(ctx *fiber.Ctx) error {
	var loginRequest model.LoginRequest

	if err := requestPkg.ValidateRequest(ctx, &loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := loginUC.ProcessLoginForm(ctx, loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return responsePkg.BuildJSONRes(ctx, data)
}
