package model

type Member struct {
	ID      int64  `json:"id" db:"id" validate:"required,number"`
	ShopID  int64  `json:"shop_id" db:"shop_id" validate:"required,number"`
	Name    string `json:"name" db:"name" validate:"required,max=255"`
	Phone   string `json:"phone" db:"phone" validate:"required,max=15,startswith=08"`
	Address string `json:"address" db:"address"`
}

type ListMemberDataResponse struct {
	Total int64
	Data  []Member
}

type MemberInsertPayload struct {
	ShopID  int64  `json:"shop_id" db:"shop_id" validate:"required,number"`
	Name    string `json:"name" db:"name" validate:"required,max=255"`
	Phone   string `json:"phone" db:"phone" validate:"required,max=15,startswith=08"`
	Address string `json:"address" db:"required"`
}

type MemberDeletePayload struct {
	ID int64 `json:"id" db:"id" validate:"required,number"`
}
