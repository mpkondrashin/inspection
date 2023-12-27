package main

import (
	"fyne.io/fyne/v2"
)

const configFileName = "config.yaml"

type Page interface {
	Name() string
	Content(win fyne.Window, model *Model) fyne.CanvasObject
	AquireData(model *Model) error
}

func main() {
	c := NewNSHIControl()
	c.Run()
}
