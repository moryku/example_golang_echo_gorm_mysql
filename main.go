package main

import (
	"github.com/heroku/go-getting-started/app/db"
	"github.com/heroku/go-getting-started/app/route"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db.InitSqlite3()
	e := route.Init()
	e.Logger.Fatal(e.Start(":" + port))
}


