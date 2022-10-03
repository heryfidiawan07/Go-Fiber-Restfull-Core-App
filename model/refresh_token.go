package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	Id        string    `gorm:"size:36;uniqueIndex;primaryKey" json:"id"`
	Revoked   bool      `gorm:"size:1;not null" json:"revoked"`
	ExpiredAt time.Time `gorm:"not null" json:"expired_at"`
	UserId    string    `gorm:"size:36;index;not null" json:"user_id"`
}

func (model *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
