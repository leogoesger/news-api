package users

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leogoesger/news-api/api/responses"
	"github.com/leogoesger/news-api/db/models"
)

// GetUser return users
func (userCtrl *Ctrl) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(userCtrl.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}