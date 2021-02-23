package models

// Category model
type Category struct {
	name string `gorm:"size:100" json: "name" form: "name"`
}
