package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string         `gorm:"unique_index;size:36;uniqueIndex;primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Username  string         `gorm:"index;not null" json:"username"`
	Email     string         `gorm:"index;not null" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	RoleId    string         `gorm:"size:36;index;" json:"role_id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (model *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
