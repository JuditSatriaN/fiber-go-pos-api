package request

import (
	"strconv"

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
			switch err.Tag() {
			case constant.ErrFieldRequired:
				return errorPkg.ConvertErrorRequired(err.Field())
			case constant.ErrFieldStartsWith:
				return errorPkg.ConvertErrorStartsWith(err.Field(), err.Param())
			case constant.ErrFieldMax:
				return errorPkg.ConvertErrorMax(err.Field(), err.Param())
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

// BuildShopIDRequest global function to build shop_id request
func BuildShopIDRequest(ctx *fiber.Ctx) (shopID int64, err error) {
	shopIDStr := ctx.Query("shop_id", "1")
	if shopIDStr != "" {
		shopID, err = strconv.ParseInt(shopIDStr, 10, 64)
		if err != nil {
			return shopID, constant.ErrShopIDMustBeNumber
		}
	}
	return shopID, nil
}

// BuildPageAndLimit global function to build page and limit
func BuildPageAndLimit(ctx *fiber.Ctx) (page int, limit int, err error) {
	page, err = BuildPage(ctx)
	if err != nil {
		return 0, 0, err
	}

	limit, err = BuildLimit(ctx)
	if err != nil {
		return 0, 0, err
	}

	return page, limit, nil
}

// BuildPage global function to build page if page == 0 will return default page
func BuildPage(ctx *fiber.Ctx) (page int, err error) {
	page, err = strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return 0, constant.ErrPageMustBeNumber
	}

	if page == 0 {
		return constant.DefaultPage, nil
	}
	return page, nil
}

// BuildLimit global function to build limit if limit == 0 will return default limit
func BuildLimit(ctx *fiber.Ctx) (limit int, err error) {
	limit, err = strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return 0, constant.ErrLimitMustBeNumber
	}

	if limit == 0 {
		return constant.DefaultLimit, nil
	}
	return limit, nil
}

// BuildOffset returns the offset from a page number
func BuildOffset(page int, limit int) int {
	if page == 0 {
		page = constant.DefaultPage
	}

	if limit == 0 {
		limit = constant.DefaultLimit
	}

	return (page * limit) - limit
}

// BuildPageLimitAndSearch global function to build page, limit, search
func BuildPageLimitAndSearch(ctx *fiber.Ctx) (page int, limit int, search string, err error) {
	page, err = BuildPage(ctx)
	if err != nil {
		return 0, 0, search, err
	}

	limit, err = BuildLimit(ctx)
	if err != nil {
		return 0, 0, search, err
	}

	search = ctx.Query("search", "")
	return page, limit, search, nil
}
