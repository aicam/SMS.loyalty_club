package handler

import (
	"github.com/aicam/scotechShops/internal/database"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func (s *Server) AddNewCustomer(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		phonenumber := vars["phonenumber"]
		db.Save(&database.CustomersData{
			PhoneNumber: phonenumber,
		})
	}
}
