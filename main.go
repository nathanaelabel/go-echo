package main

import (
	"depd_go_echo/db"
	"depd_go_echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
