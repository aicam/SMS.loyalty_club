package handler

import (
	"encoding/json"
	"github.com/aicam/scotechShops/internal/authentication"
	"github.com/aicam/scotechShops/internal/database"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func (s *Server) LoginOnce(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		username := vars["username"]
		password := vars["password"]
		user := database.ShopUsers{}
		isNull := db.Where(&database.ShopUsers{
			Username: username,
			Password: password,
		}).First(&user).RecordNotFound()
		if isNull {
			responseJson, _ := json.Marshal(LoginResponse{
				Status: 0,
				ShopId: 0,
			})
			_, _ = writer.Write(responseJson)
			return
		}
		JWT, expirationTime := authentication.GenerateJWT(username)
		responseJson, _ := json.Marshal(LoginResponse{
			Status: 1,
			ShopId: int(user.ShopId),
			JWT: JWT,
			Expiration: expirationTime,
		})
		_, _ = writer.Write(responseJson)
	}
}
