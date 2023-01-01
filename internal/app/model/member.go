package model

type Member struct {
	ID      int64  `json:"id" db:"id"`
	ShopID  int64  `json:"shop_id" db:"shop_id"`
	Name    string `json:"name" db:"name" validate:"max=255"`
	Phone   string `json:"phone" db:"phone" validate:"max=15"`
	Address string `json:"address" db:"address"`
}

type ListMemberDataResponse struct {
	Total int64
	Data  []Member
}
