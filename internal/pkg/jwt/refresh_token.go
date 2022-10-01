package jwt

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/gofiber/fiber/v2"

	responsePkg "github.com/fiber-go-pos-api/internal/pkg/response"
	jwtWare "github.com/gofiber/jwt/v3"
)

// RefreshTokenMiddleware function to handle refresh token middleware
func RefreshTokenMiddleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningMethod:  constant.JWTMethod,
		SuccessHandler: refreshTokenSuccess,
		ErrorHandler:   refreshTokenError,
		SigningKey:     GetPrivateKey().Public(),
		TokenLookup:    "header:Authorization,cookie:" + constant.JWTRefreshCookiesKey,
	})
}

func refreshTokenSuccess(ctx *fiber.Ctx) error {
	// Claim data
	// user := ctx.Locals(constant.JWTLocalsKey).(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)

	// convert all claims data
	// userID := claims["user_id"].(string)

	// if err := formsUC.RefreshTokenJWT(ctx, userID); err != nil {
	// 	return err
	// }

	return ctx.Next()
}

func refreshTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constant.ErrMissingOrMalformedJWT {
		return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusUnauthorized,
			Message:      constant.ErrMissingOrMalformedJWT,
		})
	}
	return responsePkg.BuildStandardResponse(ctx, constant.StandardResponse{
		ResponseCode: fiber.StatusUnauthorized,
		Message:      constant.ErrInvalidORExpiredJWT,
	})
}
