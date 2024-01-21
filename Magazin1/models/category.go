package models

import "time"

type CategoryModel struct {
	ID        int
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []ProductModel
}
