package db
import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitSqlite3() {
	db, err = gorm.Open("sqlite3", "./app/db/moryku_news.db")

	if err != nil {
		panic("DB Connection Error")
	}
	migrate()
}
