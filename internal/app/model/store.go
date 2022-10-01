package model

type StoreStats struct {
	StoreID        string `json:"store_id" db:"store_id" validate:"max=30"`
	TotalProduct   int64  `json:"total_product" db:"total_product"`
	TotalInventory int64  `json:"total_inventory" db:"total_inventory"`
}
