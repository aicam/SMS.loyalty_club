package handler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	db     *gorm.DB
	Router *mux.Router
}

func NewServer() *Server {
	return &Server{
		db:     nil,
		Router: mux.NewRouter(),
	}
}

func (s *Server) Routes(db *gorm.DB) {
	s.Router.HandleFunc("/login/{username}/{password}", s.LoginOnce(db))
	s.Router.HandleFunc("/add_new_user/{phonenumber}", s.AuthorizationMiddleWare(s.AddNewCustomer(db)))
}
