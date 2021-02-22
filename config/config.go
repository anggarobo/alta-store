package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB global decalaration
var DB *gorm.DB

//InitDB for connecting
func InitDB() {
	type Config struct {
		DBUsername string
		DBPassword string
		DBPort     string
		DBHost     string
		DBName     string
	}

	config := Config{
		DBUsername: "root",
		DBPassword: "",
		DBPort:     "3306",
		DBHost:     "127.0.0.1",
		DBName:     "alta_store",
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}
