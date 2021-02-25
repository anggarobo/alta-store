package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if e := config.DB.Find(&products).Error; e != nil {
		return nil, e
	}
	return products, nil
}

func GetProductByCategory(categoryID string) (interface{}, error) {
	var products models.Product
	if err := config.DB.Where("category_id =?", categoryID).Find(&products).Error; err != nil {
		return false, err
	}

	if products.ID == 0 {
		return false, nil
	}
	return products, nil
}

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
