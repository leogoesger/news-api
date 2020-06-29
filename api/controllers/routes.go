package controllers

import (
	"github.com/leogoesger/news-api/api/controllers/topics"
	"github.com/leogoesger/news-api/api/controllers/users"
	"github.com/leogoesger/news-api/api/middlewares"
)

func (s *Server) initializeRoutes() {

	_s := s.Router.PathPrefix("/api/v1").Subrouter()
	_s.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Ping)).Methods("GET")
	
	user:= users.Ctrl{DB: s.DB}
	user.CreateRoutes(_s)

	topic:= topics.Ctrl{DB: s.DB}
	topic.CreateRoutes(_s)

}

