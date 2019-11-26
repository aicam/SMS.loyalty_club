package handler

import "time"

type LoginResponse struct {
	Status int `json:"status"`
	ShopId int `json:"shop_id"`
	
	
	
	JWT string `json:"jwt"`
	Expiration time.Time `json:"expiration"`
} 
