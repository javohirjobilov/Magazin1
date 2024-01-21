package models

import "time"

type ProductModel struct {
	ID          int
	ProductType string
	Quantity    int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Products    []ProductModel
}
