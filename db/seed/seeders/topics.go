package seeders

import (
	"github.com/leogoesger/news-api/db/models"
)

// Topics seeders
var Topics = []models.Topic{
	{
		Title:   "Title 1",
		Content: "Hello world 1",
		UserID: 1,
	},
	{
		Title:   "Title 2",
		Content: "Hello world 2",
		UserID: 2,
	},
}
