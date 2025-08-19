package models

import (
	"gorm.io/gorm"
)

// Project represents the project model in the database
type Project struct {
	gorm.Model
	ProjLink     string `json:"projlink" gorm:"type:varchar(255)"`
	Title        string `json:"title" gorm:"type:varchar(255)"`
	Type         string `json:"type" gorm:"type:varchar(100)"`
	Cover        string `json:"cover" gorm:"type:varchar(255)"`
	Logo         string `json:"logo" gorm:"type:varchar(255)"`
	ProfileName  string `json:"profilename" gorm:"type:varchar(255)"`
	ProfilePic   string `json:"profilepic" gorm:"type:varchar(255)"`
	AboutProject string `json:"aboutproject" gorm:"type:text"`
	Technologies string `json:"technologies" gorm:"type:text"`
	LinkedinLink string `json:"linkedin_link" gorm:"type:varchar(255)"`
	TelegramLink string `json:"telegram_link" gorm:"type:varchar(255)"`
	XLink        string `json:"x_link" gorm:"type:varchar(255)"`
	YoutubeLink  string `json:"youtube_link" gorm:"type:varchar(255)"`
	GithubLink   string `json:"github_link" gorm:"type:varchar(255)"`
	InstaLink    string `json:"insta_link" gorm:"type:varchar(255)"`
}
