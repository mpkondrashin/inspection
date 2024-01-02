package main

import (
	"os"

	"fyne.io/fyne/v2"
)

const (
	appName        = "GInspection"
	configFileName = "inspection.yaml"
	appID          = "com.github.mpkondrashin.inspection"
)

type Page interface {
	Name() string
	Content(win fyne.Window, model *Model) fyne.CanvasObject
	AquireData(model *Model) error
}

func main() {
	capturesFolder := ""
	if len(os.Args) == 3 && os.Args[1] == "--capture" {
		capturesFolder = os.Args[2]
	}
	c := NewNSHIControl(capturesFolder)
	c.Run()
}
