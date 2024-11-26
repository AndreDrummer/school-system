package main

import (
	"school-system/cmd/app/controller"
	"school-system/cmd/app/db"
	"school-system/cmd/app/view"
)

func main() {
	db.Init()
	controller.LoadStudentsFromDB()
	view.Run()
}
