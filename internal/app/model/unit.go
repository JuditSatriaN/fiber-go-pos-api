package model

type Unit struct {
	ID     int64  `json:"id" db:"id"`
	ShopID int64  `json:"shop_id" db:"shop_id"`
	Name   string `json:"name" db:"name" validate:"max=30"`
}
