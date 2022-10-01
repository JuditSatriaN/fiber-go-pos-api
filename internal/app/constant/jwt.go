package constant

// JWT Token constants

const JWTMethod = "RS256"
const JWTLocalsKey = "user"

// Web JWT token constants

const JWTAccessCookiesKey = "jwtAccessToken"
const JWTRefreshCookiesKey = "jwtRefreshToken"

type JWTTokenKey struct {
	AccessToken  string
	RefreshToken string
}

type JWTRequest struct {
	UserID  string
	Name    string
	IsAdmin bool
}
