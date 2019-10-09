package model

import "github.com/jinzhu/gorm"

type Authentication struct {
	gorm.Model
	User   User `gorm:"foreignkey:UserID; not null"`
	UserID int
	Token  string `gorm:"type:varchar(200); not null"`
}