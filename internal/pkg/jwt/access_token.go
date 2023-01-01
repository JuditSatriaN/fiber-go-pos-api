package jwt

import (
	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/pkg/response"
	"github.com/gofiber/fiber/v2"

	jwtWare "github.com/gofiber/jwt/v3"
)

// AccessTokenMiddleware function to handle access token middleware
func AccessTokenMiddleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		ErrorHandler:  accessTokenError,
		SigningMethod: constant.JWTMethod,
		SigningKey:    GetPrivateKey().Public(),
		TokenLookup:   "header:Authorization,cookie:" + constant.JWTAccessCookiesKey,
	})
}

// accessTokenError custom package to handle jwt access token error
func accessTokenError(ctx *fiber.Ctx, err error) error {
	if err.Error() == constant.ErrMissingOrMalformedJWT {
		return response.BuildStandardResponse(ctx, constant.StandardResponse{
			ResponseCode: fiber.StatusUnauthorized,
			Message:      constant.ErrMissingOrMalformedJWT,
		})
	}

	return ctx.Next()
}
