package model

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name" validate:"required,max=30"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	ID              int64  `json:"id" form:"id"`
	UserName        string `json:"user_name" form:"user_name"`
	FullName        string `json:"full_name" form:"full_name"`
	IsAdmin         bool   `json:"is_admin" form:"is_admin"`
	JWTAccessToken  string `json:"jwt_access_token" form:"jwt_access_token"`
	JWTRefreshToken string `json:"jwt_refresh_token" form:"jwt_refresh_token"`
}
