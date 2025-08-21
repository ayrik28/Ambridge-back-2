package models

import (
	"gorm.io/gorm"
)

// Crew represents the crew member model in the database
type Crew struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(100)"`
	About    string `json:"about" gorm:"type:text"`
	URLPhoto string `json:"urlphoto" gorm:"type:varchar(255)"`
}
