package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Votes struct for db
type Votes struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Score     float32   `gorm:"type:double precision;default:0" json:"score"`
	User      User      `json:"user"`
	News      News      `json:"topic"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	NewsID    uint32    `gorm:"not null" json:"topic_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare news to insert
func (p *Votes) Prepare() {
	p.ID = 0
	p.User = User{}
	p.Score = 0
	p.News = News{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate news
func (p *Votes) Validate() error {
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	if p.NewsID < 1 {
		return errors.New("Required News")
	}
	return nil
}

// SaveVotes create news
func (p *Votes) SaveVotes(db *gorm.DB) (*Votes, error) {
	var err error
	err = db.Debug().Model(&Votes{}).Create(&p).Error
	if err != nil {
		return &Votes{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Votes{}, err
		}
	}
	return p, nil
}

// FindAllVotess get all newss
func (p *Votes) FindAllVotess(db *gorm.DB) (*[]Votes, error) {
	var err error
	Votess := []Votes{}
	err = db.Debug().Model(&Votes{}).Limit(100).Find(&Votess).Error
	if err != nil {
		return &[]Votes{}, err
	}
	if len(Votess) > 0 {
		for i := range Votess {
			err := db.Debug().Model(&User{}).Where("id = ?", Votess[i].UserID).Take(&Votess[i].User).Error
			if err != nil {
				return &[]Votes{}, err
			}
		}
	}
	return &Votess, nil
}

// FindVotesByID get news by id
func (p *Votes) FindVotesByID(db *gorm.DB, pid uint64) (*Votes, error) {
	var err error
	err = db.Debug().Model(&Votes{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Votes{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Votes{}, err
		}
	}
	return p, nil
}

// UpdateAVotes update a news
func (p *Votes) UpdateAVotes(db *gorm.DB) (*Votes, error) {

	var err error

	err = db.Debug().Model(&Votes{}).Where("id = ?", p.ID).Updates(Votes{Score: p.Score, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Votes{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Votes{}, err
		}
	}
	return p, nil
}

// DeleteAVotes delete a news
func (p *Votes) DeleteAVotes(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Votes{}).Where("id = ? and User_id = ?", pid, uid).Take(&Votes{}).Delete(&Votes{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Votes not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}