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

func CreateCategory(category *models.Category) (interface{}, error) {
	if err := config.DB.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
