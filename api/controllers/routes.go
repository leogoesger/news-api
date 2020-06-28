package controllers

import "github.com/leogoesger/news-api/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	_s := s.Router.PathPrefix("/api/v1").Subrouter()
	_s.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	_s.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	_s.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	_s.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	_s.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	_s.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	_s.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Topics routes
	_s.HandleFunc("/topics", middlewares.SetMiddlewareJSON(s.CreateTopic)).Methods("POST")
	_s.HandleFunc("/topics", middlewares.SetMiddlewareJSON(s.GetTopics)).Methods("GET")
	_s.HandleFunc("/topics/{id}", middlewares.SetMiddlewareJSON(s.GetTopic)).Methods("GET")
	_s.HandleFunc("/topics/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateTopic))).Methods("PUT")
	_s.HandleFunc("/topics/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteTopic)).Methods("DELETE")
}