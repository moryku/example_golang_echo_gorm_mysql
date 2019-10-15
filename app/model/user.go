package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"not null form:"email"; unique"`
	Password  string `gorm:"not null form:"password"" json:"-"`
	Name      string `gorm:"not null form:"name"; type:varchar(100)"` // unique_index
}

