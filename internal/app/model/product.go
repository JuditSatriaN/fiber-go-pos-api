package model

type Product struct {
	ID                 int64   `json:"id" db:"id" validate:"required"`
	ShopID             int64   `json:"shop_id" db:"shop_id" validate:"required"`
	UnitID             int64   `json:"unit_id" db:"unit_id"`
	Name               string  `json:"name" db:"name" validate:"max=255"`
	Stock              int64   `json:"stock" db:"stock"`
	Barcode            string  `json:"barcode" db:"barcode" validate:"max=15"`
	Price              float64 `json:"price" db:"price"`
	MemberPrice        float64 `json:"member_price" db:"member_price"`
	Discount           float64 `json:"discount" db:"discount"`
	DiscountPercentage int32   `json:"discount_percentage" db:"discount_percentage"`
	Purchase           float64 `json:"purchase" db:"purchase"`
	Ppn                bool    `json:"ppn" db:"ppn"`
}

type ListProductDataResponse struct {
	Total int64
	Data  []Product
}
