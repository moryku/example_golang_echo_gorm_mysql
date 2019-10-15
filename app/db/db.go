package db

import (
	"github.com/heroku/go-getting-started/app/config"
	"github.com/heroku/go-getting-started/app/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_HOST,
		configuration.DB_PORT,
		configuration.DB_NAME)
	fmt.Println(connect_string)
	db, err = gorm.Open("mysql", connect_string)

	if err != nil {
		panic("DB Connection Error")
	}
	migrate()
}

func migrate() {
	db.Model(&model.User{}).DropColumn("address")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Authentication{})
	db.Model(&model.Authentication{}).AddForeignKey("user_id",
		"users(id)", "CASCADE", "CASCADE")
}

func DbManager() *gorm.DB {
	return db
}
