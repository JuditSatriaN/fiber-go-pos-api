package model

type Product struct {
	PLU     string `json:"plu" db:"plu" validate:"required,max=30"`
	Name    string `json:"name" db:"name" validate:"max=255"`
	Barcode string `json:"barcode" db:"barcode" validate:"max=30"`
	Ppn     bool   `json:"ppn" db:"ppn"`
}

type ListProductDataResponse struct {
	Total int64
	Data  []Product
}
