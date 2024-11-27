package main

import (
	schoolsystem "school-system/cmd/app/controller"
	"school-system/cmd/app/db"
	"school-system/cmd/app/view"
)

func main() {
	db.Init()
	schoolsystem.Init()
	view.Run()
}
