package member

import (
	"github.com/gofiber/fiber/v2"

	memberHandler "github.com/fiber-go-pos-api/internal/app/handler/member"
)

// BuildMemberAPI : API to handle member
func BuildMemberAPI(api fiber.Router) {
	api.Get("/members", memberHandler.GetAllMemberHandler)
	api.Post("/member/insert", memberHandler.InsertMemberHandler)
	api.Post("/member/update", memberHandler.UpdateMemberHandler)
	api.Post("/member/delete", memberHandler.DeleteMemberHandler)
	api.Post("/member/upsert", memberHandler.UpsertMemberHandler)
}
