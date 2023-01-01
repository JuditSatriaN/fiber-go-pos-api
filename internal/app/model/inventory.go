package model

type Inventory struct {
	ID                 int64   `json:"id" db:"id"`
	ShopID             int64   `json:"shop_id" db:"shop_id"`
	PLU                string  `json:"plu" db:"plu" validate:"max=30"`
	Name               string  `json:"name" db:"name" validate:"max=255"`
	UnitID             int64   `json:"unit_id" db:"unit_id"`
	UnitName           string  `json:"unit_name" db:"unit_name" validate:"max=30"`
	Barcode            string  `json:"barcode" db:"barcode" validate:"max=30"`
	Ppn                bool    `json:"ppn" db:"ppn"`
	Multiplier         int64   `json:"multiplier" db:"multiplier"`
	Stock              int64   `json:"stock" db:"stock"`
	Price              float64 `json:"price" db:"price"`
	MemberPrice        float64 `json:"member_price" db:"member_price"`
	Purchase           float64 `json:"purchase" db:"purchase"`
	Discount           float64 `json:"discount" db:"discount"`
	DiscountPercentage int32   `json:"discount_percentage" db:"discount_percentage" validate:"max=100"`
}

type ListInventoryDataResponse struct {
	Total int64
	Data  []Inventory
}

type UpdateStockAfterSalesData struct {
	ID  int64 `json:"id" db:"id"`
	Qty int32 `json:"qty" db:"qty"`
}
