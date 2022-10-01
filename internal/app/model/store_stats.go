package model

type StoreStats struct {
	StoreID        int64 `json:"store_id" db:"store_id"`
	TotalProduct   int64 `json:"total_product" db:"total_product"`
	TotalInventory int64 `json:"total_inventory" db:"total_inventory"`
}
