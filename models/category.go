package models

import (
	"gorm.io/gorm"
)

// Category model
type Category struct {
	gorm.Model
	Name string `gorm:"size:100" json: "name" form: "name"`
}