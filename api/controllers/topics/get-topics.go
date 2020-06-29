package topics

import (
	"net/http"

	"github.com/leogoesger/news-api/api/responses"
	"github.com/leogoesger/news-api/db/models"
)


// GetTopics get topics
func (topicCtrl *Ctrl) GetTopics(w http.ResponseWriter, r *http.Request) {

	post := models.Topic{}

	posts, err := post.FindAllTopics(topicCtrl.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}