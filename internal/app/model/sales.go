package model

type Sales struct {
	Head   SalesHead     `json:"sales_head"`
	Detail []SalesDetail `json:"sales_detail"`
}

type SalesHead struct {
	Invoice       string  `json:"invoice" db:"invoice" validate:"max=15"`
	UserID        string  `json:"user_id" db:"user_id" validate:"max=30"`
	TotalItem     int     `json:"total_item" db:"total_item"`
	TotalPrice    float32 `json:"total_price" db:"total_price"`
	TotalPurchase float32 `json:"total_purchase" db:"total_purchase"`
	TotalTax      float32 `json:"total_tax" db:"total_tax"`
	TotalDiscount float32 `json:"total_discount" db:"total_discount"`
	TotalPay      float32 `json:"total_pay" db:"total_pay"`
}

type SalesDetail struct {
	Invoice     string  `json:"invoice" db:"invoice" validate:"max=15"`
	UserID      string  `json:"user_id" db:"user_id" validate:"max=30"`
	PLU         string  `json:"plu" db:"plu" validate:"max=30"`
	Name        string  `json:"name" db:"name" validate:"max=255"`
	UnitName    string  `json:"unit_name" db:"unit_name" validate:"max=30"`
	Barcode     string  `json:"barcode" db:"barcode" validate:"max=30"`
	Ppn         bool    `json:"ppn" db:"ppn"`
	Qty         int64   `json:"qty" db:"qty"`
	Price       float32 `json:"price" db:"price"`
	Purchase    float32 `json:"purchase" db:"purchase"`
	Discount    float32 `json:"discount" db:"discount"`
	MemberID    int     `json:"member_id" db:"member_id"`
	InventoryID int32   `json:"inventory_id" db:"inventory_id"`
}
