package pagination

import (
	"strconv"

	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

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
