package database

import (
	"alta-store/config"
	"alta-store/middlewares"
	"alta-store/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUsers(user *models.Users) (interface{}, error) {

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUsers(user *models.Users) (interface{}, error) {
	if err := config.DB.Where("email = ? AND password = ?",
		user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	token, err := middlewares.CreateToken(int(1))
	if err != nil {
		return nil, err
	}
	user.Token = token

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.Users
	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}
