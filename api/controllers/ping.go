package controllers

import (
	"net/http"

	"github.com/leogoesger/news-api/api/responses"
)

// Ping controller
func (server *Server) Ping(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}