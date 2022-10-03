package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	Id        string `gorm:"unique_index;size:36;uniqueIndex;primaryKey" json:"id"`
	ParenMenu string `json:"parent_menu"`
	ParentId  string `json:"parent_id"`
	Name      string `gorm:"unique_index;not null" json:"name"`
	Alias     string `json:"alias"`
	Url       string `json:"url"`
	Icon      string `json:"icon"`
}

func (model *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
