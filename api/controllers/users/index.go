package users

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/leogoesger/news-api/api/middlewares"
)

// Ctrl controller struct
type Ctrl struct {
	DB     *gorm.DB
}

// CreateRoutes generate routes
func (userCtrl *Ctrl) CreateRoutes (s *mux.Router){
	_s := s.PathPrefix("/users").Subrouter()

	_s.HandleFunc("/login", middlewares.SetMiddlewareAuthentication(userCtrl.Login)).Methods("POST")
	_s.HandleFunc("/", middlewares.SetMiddlewareJSON(userCtrl.CreateUser)).Methods("POST")
	_s.HandleFunc("/", middlewares.SetMiddlewareJSON(userCtrl.GetUsers)).Methods("GET")
	_s.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(userCtrl.GetUser)).Methods("GET")
	_s.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(userCtrl.UpdateUser))).Methods("PUT")
	_s.HandleFunc("/{id}", middlewares.SetMiddlewareAuthentication(userCtrl.DeleteUser)).Methods("DELETE")
}