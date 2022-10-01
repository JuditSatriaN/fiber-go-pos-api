package model

type User struct {
	ID       int64  `json:"id" db:"id" validate:"required"`
	ShopID   int64  `json:"shop_id" db:"shop_id"`
	UserName string `json:"user_name" db:"user_name" validate:"max=30"`
	FullName string `json:"full_name" db:"full_name" validate:"max=255"`
	Password string `json:"password,omitempty" db:"password" validate:"max=255"`
	IsAdmin  bool   `json:"is_admin,omitempty" db:"is_admin"`
}

type ListUserDataResponse struct {
	Total int64
	Data  []User
}
