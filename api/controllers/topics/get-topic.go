package topics

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leogoesger/news-api/api/responses"
	"github.com/leogoesger/news-api/db/models"
)


// GetTopic get topic
func (topicCtrl *Ctrl) GetTopic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post := models.Topic{}

	postReceived, err := post.FindTopicByID(topicCtrl.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, postReceived)
}