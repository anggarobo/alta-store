package models

import (
	"gorm.io/gorm"
)

// Category model
type Category struct {
	gorm.Model
	Name string `gorm:"size:100" json: "name" form: "name"`
}

// func (Category *Category) TableName() string {
// 	return "categories"
// }

// func (Category *Category) toString() string {
// 	return fmt.Sprintf("id: %d\nname: %s", Category.ID, Category.Name)
// }
