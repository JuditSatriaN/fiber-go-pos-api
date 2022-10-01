package model

type Sales struct {
	Head   SalesHead     `json:"sales_head"`
	Detail []SalesDetail `json:"sales_detail"`
}

type SalesHead struct {
	ID            int64   `json:"id" db:"id"`
	ShopID        int64   `json:"shop_id" db:"shop_id"`
	Invoice       string  `json:"invoice" db:"invoice" validate:"max=15"`
	UserID        int64   `json:"user_id" db:"user_id"`
	TotalItem     int32   `json:"total_item" db:"total_item"`
	TotalPrice    float32 `json:"total_price" db:"total_price"`
	TotalPurchase float32 `json:"total_purchase" db:"total_purchase"`
	TotalTax      float32 `json:"total_tax" db:"total_tax"`
	TotalDiscount float32 `json:"total_discount" db:"total_discount"`
	TotalPay      float32 `json:"total_pay" db:"total_pay"`
}

type SalesDetail struct {
	ID          int64   `json:"id" db:"id"`
	ShopID      int64   `json:"shop_id" db:"shop_id"`
	Invoice     string  `json:"invoice" db:"invoice" validate:"max=30"`
	UserID      int64   `json:"user_id" db:"user_id"`
	ProductID   int64   `json:"product_id" db:"product_id"`
	ProductName string  `json:"product_name" db:"product_name" validate:"max=255"`
	UnitID      int64   `json:"unit_id" db:"unit_id"`
	UnitName    string  `json:"unit_name" db:"unit_name" validate:"max=30"`
	Barcode     string  `json:"barcode" db:"barcode" validate:"max=15"`
	Ppn         bool    `json:"ppn" db:"ppn"`
	Qty         int32   `json:"qty" db:"qty"`
	Price       float32 `json:"price" db:"price"`
	Purchase    float32 `json:"purchase" db:"purchase"`
	Discount    float32 `json:"discount" db:"discount"`
	MemberID    int     `json:"member_id" db:"member_id"`
}
