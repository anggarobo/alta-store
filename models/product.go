package models

import (
	"time"
)

// Product Model
type Product struct {
	ID          int
	Name        string
	Stock       int
	Price       int
	Description string
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	ID   int
	Name string
}
