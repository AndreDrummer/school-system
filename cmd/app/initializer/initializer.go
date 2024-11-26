package initializer

import (
	"school-system/cmd/app/db"
)

func initSystem() {
	db.Init()

	loadStudentsFromDB()
}

func Run() {
	initSystem()
}
