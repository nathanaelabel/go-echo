package main

import "depd_go_echo/routes"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
