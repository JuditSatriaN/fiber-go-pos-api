package request

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	errorPkg "github.com/fiber-go-pos-api/internal/pkg/error"
)

// ValidateRequest global function to validate request
func ValidateRequest(ctx *fiber.Ctx, dest any) error {
	if err := ctx.BodyParser(dest); err != nil {
		return err
	}

	if err := validator.New().Struct(dest); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == constant.ErrFieldStartsWith {
				return errorPkg.ConvertErrorStartsWith(err.Field(), err.Param())
			}
		}
		return err
	}

	return nil
}

// BuildSearchRequest global function to build search request
func BuildSearchRequest(ctx *fiber.Ctx) string {
	return ctx.Query("search", "")
}
