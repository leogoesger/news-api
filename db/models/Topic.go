package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Topic struct for db
type Topic struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Score     float32   `gorm:"type:double precision;default:0" json:"score"`
	User      User      `json:"user"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare topic to insert
func (p *Topic) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.User = User{}
	p.Score = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate topic
func (p *Topic) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

// SaveTopic create topic
func (p *Topic) SaveTopic(db *gorm.DB) (*Topic, error) {
	var err error
	err = db.Debug().Model(&Topic{}).Create(&p).Error
	if err != nil {
		return &Topic{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Topic{}, err
		}
	}
	return p, nil
}

// FindAllTopics get all topics
func (p *Topic) FindAllTopics(db *gorm.DB) (*[]Topic, error) {
	var err error
	Topics := []Topic{}
	err = db.Debug().Model(&Topic{}).Limit(100).Find(&Topics).Error
	if err != nil {
		return &[]Topic{}, err
	}
	if len(Topics) > 0 {
		for i := range Topics {
			err := db.Debug().Model(&User{}).Where("id = ?", Topics[i].UserID).Take(&Topics[i].User).Error
			if err != nil {
				return &[]Topic{}, err
			}
		}
	}
	return &Topics, nil
}

// FindTopicByID get topic by id
func (p *Topic) FindTopicByID(db *gorm.DB, pid uint64) (*Topic, error) {
	var err error
	err = db.Debug().Model(&Topic{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Topic{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Topic{}, err
		}
	}
	return p, nil
}

// UpdateATopic update a topic
func (p *Topic) UpdateATopic(db *gorm.DB) (*Topic, error) {

	var err error

	err = db.Debug().Model(&Topic{}).Where("id = ?", p.ID).Updates(Topic{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Topic{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Topic{}, err
		}
	}
	return p, nil
}

// DeleteATopic delete a topic
func (p *Topic) DeleteATopic(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Topic{}).Where("id = ? and User_id = ?", pid, uid).Take(&Topic{}).Delete(&Topic{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Topic not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}