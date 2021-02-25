package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetCategories() (interface{}, error) {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategory(id int) (interface{}, error) {
	var category models.Category

	if err := config.DB.Find(&category, id).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func CreateCategory(category *models.Category) (interface{}, error) {
	if err := config.DB.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
