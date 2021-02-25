package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetCarts(cartID int) (interface{}, error) {
	var cartDetail []models.Cart_details
	if e := config.DB.Where("cart_id = ? ", cartID).Find(&cartDetail).Error; e != nil {
		return nil, e
	}
	return cartDetail, nil
}

func GetCartId(userID int) (int, error) {
	var cart models.Carts
	if e := config.DB.Where("user_id = ? AND status = 0", userID).Find(&cart).Error; e != nil {
		return 0, e
	}

	ID := int(cart.ID)
	if ID == 0 {
		cart.User_id = userID
		if e := config.DB.Create(&cart).Error; e != nil {
			return 0, e
		}
	}
	return int(cart.ID), nil
}

func GetProductPrice(productID int, quantity int) (int, error) {
	var product models.Products
	if e := config.DB.Where("ID = ?", productID).Find(&product).Error; e != nil {
		return 0, e
	}

	if product.Stock < quantity {
		return 0, nil
	}

	product.Stock = product.Stock - quantity
	if e := config.DB.Save(product).Error; e != nil {
		return 0, e
	}

	return product.Price, nil
}

func AddToCart(cartDetail models.Cart_details) (interface{}, error) {
	if e := config.DB.Create(&cartDetail).Error; e != nil {
		return nil, e
	}

	return cartDetail, nil
}

func GetCartDetailID(cartDetailID string) (bool, error) {
	var cartDetail models.Cart_details
	if e := config.DB.Where("ID = ?", cartDetailID).Find(&cartDetail).Error; e != nil {
		return false, e
	}

	if cartDetail.ID == 0 {
		return false, nil
	}

	return true, nil
}

func DeleteCart(cartDetailID string) (interface{}, error) {
	var cartDetail models.Cart_details
	if e := config.DB.Where("ID = ?", cartDetailID).Find(&cartDetail).Error; e != nil {
		return nil, e
	}

	if e := config.DB.Delete(&cartDetail, cartDetailID).Error; e != nil {
		return nil, e
	}

	var product models.Products
	if e := config.DB.Where("ID = ?", cartDetail.Product_id).Find(&product).Error; e != nil {
		return nil, e
	}

	product.Stock = product.Stock + cartDetail.Quantity
	if e := config.DB.Save(product).Error; e != nil {
		return nil, e
	}

	return cartDetailID, nil
}
