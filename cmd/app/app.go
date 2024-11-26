package app

import (
	"school-system/cmd/app/controller"
	"school-system/cmd/app/db"
	"school-system/cmd/app/view"
)

func Run() {
	db.Init()
	controller.LoadStudentsFromDB()
	view.Run()
}
