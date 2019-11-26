package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

func (s *Server) SendSmsApi(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
