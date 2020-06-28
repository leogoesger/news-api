package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/leogoesger/news-api/db/models"
	"github.com/leogoesger/news-api/db/seed/seeders"
)


// Load db
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Votes{}, &models.News{}, &models.Topic{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Topic{}, &models.News{}, &models.Votes{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Topic{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.News{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.News{}).AddForeignKey("topic_id", "topics(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Votes{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Votes{}).AddForeignKey("news_id", "news(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range seeders.Users {
		err = db.Debug().Model(&models.User{}).Create(&seeders.Users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i := range seeders.Topics {
		err = db.Debug().Model(&models.Topic{}).Create(&seeders.Topics[i]).Error
		if err != nil {
			log.Fatalf("cannot seed topics table: %v", err)
		}
	}
}