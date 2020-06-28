package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// News struct for db
type News struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Link   	  string    `gorm:"size:255;not null;" json:"link"`
	Score     float32   `gorm:"type:double precision;default:0" json:"score"`
	User      User      `json:"user"`
	Topic     Topic     `json:"topic"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	TopicID   uint32    `gorm:"not null" json:"topic_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare news to insert
func (p *News) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Link = html.EscapeString(strings.TrimSpace(p.Link))
	p.User = User{}
	p.Score = 0
	p.Topic = Topic{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate news
func (p *News) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Link == "" {
		return errors.New("Required Link")
	}
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

// SaveNews create news
func (p *News) SaveNews(db *gorm.DB) (*News, error) {
	var err error
	err = db.Debug().Model(&News{}).Create(&p).Error
	if err != nil {
		return &News{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &News{}, err
		}
	}
	return p, nil
}

// FindAllNewss get all newss
func (p *News) FindAllNewss(db *gorm.DB) (*[]News, error) {
	var err error
	Newss := []News{}
	err = db.Debug().Model(&News{}).Limit(100).Find(&Newss).Error
	if err != nil {
		return &[]News{}, err
	}
	if len(Newss) > 0 {
		for i := range Newss {
			err := db.Debug().Model(&User{}).Where("id = ?", Newss[i].UserID).Take(&Newss[i].User).Error
			if err != nil {
				return &[]News{}, err
			}
		}
	}
	return &Newss, nil
}

// FindNewsByID get news by id
func (p *News) FindNewsByID(db *gorm.DB, pid uint64) (*News, error) {
	var err error
	err = db.Debug().Model(&News{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &News{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &News{}, err
		}
	}
	return p, nil
}

// UpdateANews update a news
func (p *News) UpdateANews(db *gorm.DB) (*News, error) {

	var err error

	err = db.Debug().Model(&News{}).Where("id = ?", p.ID).Updates(News{Title: p.Title, Link: p.Link, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &News{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &News{}, err
		}
	}
	return p, nil
}

// DeleteANews delete a news
func (p *News) DeleteANews(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&News{}).Where("id = ? and User_id = ?", pid, uid).Take(&News{}).Delete(&News{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("News not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}