package database

import "github.com/jinzhu/gorm"

type ShopUsers struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	ShopId uint `json:"shop_id"`
}

type CustomersData struct {
	gorm.Model
	PhoneNumber string `json:"phone_number"`
	Score uint `json:"score"`
}