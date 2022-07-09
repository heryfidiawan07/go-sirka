package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id   string `gorm:"size:36;uniqueIndex;primaryKey"`
	Name string `gorm:"size:100;"`
}

func (model *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
