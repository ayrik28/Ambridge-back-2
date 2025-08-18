package models

import (
	"gorm.io/gorm"
)

// User represents the user model in the database
type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"type:varchar(100)"`
	Surname        string `json:"surname" gorm:"type:varchar(100)"`
	Email          string `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Password       string `json:"-" gorm:"type:varchar(255)"` // Password is not exposed in JSON
	ProfileImage   string `json:"profile_image,omitempty" gorm:"type:varchar(255)"`
	CompanyName    string `json:"company_name,omitempty" gorm:"type:varchar(100)"`
	CompanyEmail   string `json:"company_email,omitempty" gorm:"type:varchar(255)"`
	CompanyAddress string `json:"company_address,omitempty" gorm:"type:text"`
	CompanyPhone   string `json:"company_phone,omitempty" gorm:"type:varchar(20)"`
	Position       string `json:"position,omitempty" gorm:"type:varchar(100)"`
	ReferralSource string `json:"referral_source,omitempty" gorm:"type:varchar(100)"`
	Role           string `json:"role" gorm:"type:varchar(20);default:'user'"` // 'admin' or 'user'
	RefreshToken   string `json:"-" gorm:"type:varchar(255)"`
	ResumeFile     string `json:"resume_file,omitempty" gorm:"type:varchar(255)"`
}
