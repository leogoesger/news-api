package controllers

import (
	"net/http"

	"github.com/leogoesger/news-api/api/responses"
)

// Home controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}