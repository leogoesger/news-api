package topics

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
func (topicCtrl *Ctrl) CreateRoutes (s *mux.Router){
	_s := s.PathPrefix("/topics").Subrouter()

	_s.HandleFunc("/", middlewares.SetMiddlewareJSON(topicCtrl.CreateTopic)).Methods("POST")
	_s.HandleFunc("/", middlewares.SetMiddlewareJSON(topicCtrl.GetTopics)).Methods("GET")
	_s.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(topicCtrl.GetTopic)).Methods("GET")
	_s.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(topicCtrl.UpdateTopic))).Methods("PUT")
	_s.HandleFunc("/{id}", middlewares.SetMiddlewareAuthentication(topicCtrl.DeleteTopic)).Methods("DELETE")
}