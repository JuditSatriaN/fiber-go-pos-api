package model

type User struct {
	ID       int64  `json:"id" db:"id" validate:"required"`
	ShopID   int64  `json:"shop_id" db:"shop_id" validate:"required"`
	UserName string `json:"user_name" db:"user_name" validate:"required,max=30"`
	FullName string `json:"full_name" db:"full_name" validate:"required,max=255"`
	Password string `json:"password" db:"password" validate:"required,max=255"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin" validate:"required,boolean"`
}

type ListUserDataResponse struct {
	Total int64
	Data  []User
}

type UserInsertPayload struct {
	ShopID   int64  `json:"shop_id" validate:"required,number"`
	UserName string `json:"user_name" validate:"required,max=30"`
	FullName string `json:"full_name" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=255"`
	IsAdmin  bool   `json:"is_admin" validate:"required,boolean"`
}

type UserDeletePayload struct {
	ID int64 `json:"id" db:"id" validate:"required"`
}

type UserUpsertPayload struct {
	ID       int64  `json:"id" db:"id"`
	ShopID   int64  `json:"shop_id" validate:"required,number"`
	UserName string `json:"user_name" validate:"required,max=30"`
	FullName string `json:"full_name" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=255"`
	IsAdmin  bool   `json:"is_admin" validate:"required,boolean"`
}
