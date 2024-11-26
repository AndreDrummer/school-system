package app

import (
	"school-system/cmd/app/initializer"
	"school-system/cmd/app/view"
)

func Run() {
	initializer.Run()
	view.Run()
}
