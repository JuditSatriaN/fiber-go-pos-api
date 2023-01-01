package model

import "time"

type ProductSalesStatsDaily struct {
	ID        int64     `json:"id" db:"id"`
	ShopID    int64     `json:"shop_id" db:"shop_id"`
	DateSold  time.Time `json:"date_sold" db:"date_sold"`
	ProductID int64     `json:"product_id" db:"product_id"`
	TotalSold int64     `json:"total_sold" db:"total_sold"`
}
