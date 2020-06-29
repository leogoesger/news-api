package users

import (
	"net/http"

	"github.com/leogoesger/news-api/api/responses"
	"github.com/leogoesger/news-api/db/models"
)

// GetUsers return users
func (userCtrl *Ctrl) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(userCtrl.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}