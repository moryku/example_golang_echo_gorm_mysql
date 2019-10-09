package main

import (
	"simple-rest-api-moryku/app/db"
	"simple-rest-api-moryku/app/route"
)

func main() {
	db.Init()
	e := route.Init()
	e.Logger.Fatal(e.Start(":9000"))
}
