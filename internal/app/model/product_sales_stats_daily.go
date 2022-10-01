package model

import "time"

type ProductSalesStatsDaily struct {
	ShopID    int64     `json:"shop_id" db:"shop_id"`
	DateSold  time.Time `json:"date_sold" db:"date_sold"`
	PLU       string    `json:"plu" db:"plu"  validate:"max=30"`
	Name      string    `json:"name" db:"name"  validate:"max=255"`
	TotalSold int64     `json:"total_sold" db:"total_sold"`
}
