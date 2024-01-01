package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type NSHIControl struct {
	pages   []Page
	current int
	app     fyne.App
	win     fyne.Window
	model   Model
}

//go:generate fyne bundle --name IconSVGResource --output resource.go   ../../bin/icon.svg

func (c *NSHIControl) Window(p Page) fyne.CanvasObject {
	left := container.NewVBox()
	image := canvas.NewImageFromResource(IconSVGResource)
	image.SetMinSize(fyne.NewSize(52, 52))
	image.FillMode = canvas.ImageFillContain
	left.Add(image)
	for _, page := range c.pages {
		if page == p {
			left.Add(widget.NewLabelWithStyle("▶ "+page.Name(), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		} else {
			left.Add(widget.NewLabel("    " + page.Name()))
		}
	}

	middle := container.NewPadded(container.NewVBox(layout.NewSpacer(), p.Content(c.win, &c.model), layout.NewSpacer()))

	upper := container.NewBorder(nil, nil, container.NewHBox(left, widget.NewSeparator()), nil, middle)
	quitButton := widget.NewButtonWithIcon("Quit", theme.CancelIcon(), c.Quit)
	prevButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), c.Prev)
	if c.current == 0 {
		prevButton.Disable()
	}

	nextButton := widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), c.Next)
	nextButton.IconPlacement = widget.ButtonIconTrailingText

	if c.current == len(c.pages)-1 {
		nextButton.Disable()
	}

	buttons := container.NewBorder(nil, nil, quitButton,
		container.NewHBox(prevButton, nextButton))
	bottom := container.NewVBox(widget.NewSeparator(), buttons)
	_ = bottom

	return container.NewBorder(nil, container.NewPadded(bottom), nil, nil, upper)
}

func (c *NSHIControl) Quit() {
	c.app.Quit()
}

func (c *NSHIControl) Next() {
	err := c.pages[c.current].AquireData(&c.model)
	if err != nil {
		dialog.ShowError(err, c.win)
		return
	}
	c.current++
	c.win.SetContent(c.Window(c.pages[c.current]))
}
func (c *NSHIControl) Prev() {
	c.current--
	c.win.SetContent(c.Window(c.pages[c.current]))
}

func NewNSHIControl() *NSHIControl {
	c := &NSHIControl{
		app: app.New(),
		model: Model{
			appName:  appName,
			fileName: configFileName,
		},
	}
	c.win = c.app.NewWindow("Network Security Inspect/Bypass Switch")
	c.win.Resize(fyne.NewSize(600, 400))
	//c.win.SetFixedSize(true)
	c.win.SetMaster()
	c.pages = []Page{
		&PageIntro{},
		&PagePassword{},
		&PageOptions{},
		&PageRegion{},
		&PageControl{},
	}
	c.win.SetContent(c.Window(c.pages[0]))
	return c
}

func (c *NSHIControl) Run() {
	c.win.ShowAndRun()
}
