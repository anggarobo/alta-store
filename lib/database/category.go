package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetCategories() (interface{}, error) {
	var categories []models.Category

	if e := config.DB.Find(&categories).Error; e != nil {
		return nil, e
	}
	return categories, nil
}
