package handler

import (
	"github.com/aicam/scotechShops/internal/authentication"
	"net/http"
)

func (s *Server) AuthorizationMiddleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		if token == "" {
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		authorizationStatus := authentication.CheckToken(token)
		if authorizationStatus == -1 {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if authorizationStatus == 0 {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(writer, request)
	}
}
