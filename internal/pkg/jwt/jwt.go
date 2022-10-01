package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey *rsa.PrivateKey

// GenerateJWT function to generate private key jwt
func GenerateJWT(embedPrivatePEMFile []byte) error {
	var err error

	privateKeyDecode, _ := pem.Decode(embedPrivatePEMFile)
	privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyDecode.Bytes)
	return err
}

// GetPrivateKey function to get private key jwt
func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}

// CreateJWTToken function to create jwt token
func CreateJWTToken(req constant.JWTRequest) (constant.JWTTokenKey, error) {
	var (
		err         error
		jwtTokenKey constant.JWTTokenKey
	)

	// Create access token
	accessTokenClaims := jwt.MapClaims{
		"user_id":  req.UserID,
		"name":     req.Name,
		"is_admin": req.IsAdmin,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims)
	jwtTokenKey.AccessToken, err = accessToken.SignedString(privateKey)
	if err != nil {
		return constant.JWTTokenKey{}, err
	}

	// Create refresh token
	refreshTokenClaims := jwt.MapClaims{
		"user_id": req.UserID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims)
	jwtTokenKey.RefreshToken, err = refreshToken.SignedString(privateKey)
	if err != nil {
		return constant.JWTTokenKey{}, err
	}

	return jwtTokenKey, nil
}
