package model

type Store struct {
	ID      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name" validate:"max=50"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone" validate:"max=15"`
}
