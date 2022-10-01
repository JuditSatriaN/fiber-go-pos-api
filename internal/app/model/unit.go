package model

type Unit struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" validate:"max=30"`
}
