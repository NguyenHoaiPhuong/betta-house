package main

import "github.com/NguyenHoaiPhuong/betta-house/app"

func main() {
	app := new(app.App)
	app.Init()
	app.PrintConfig()
	app.Run()
}
