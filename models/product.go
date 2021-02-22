package models

import "time"

// Product Model
type Product struct {
	name        string
	stock       int
	price       int
	description string
	createdAt   time.Time
	updatedAt   time.Time
}
