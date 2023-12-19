package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NSHIControl struct {
	pages   []Page
	current int
	app     fyne.App
	win     fyne.Window
	model   Model
}

func (c *NSHIControl) Window(p Page) fyne.CanvasObject {
	left := container.NewVBox()
	for _, page := range c.pages {
		if page == p {
			left.Add(widget.NewLabelWithStyle("▶ "+page.Name(), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		} else {
			left.Add(widget.NewLabel("    " + page.Name()))
		}
	}

	middle := container.NewPadded(container.NewVBox(layout.NewSpacer(), p.Content(c.win, &c.model), layout.NewSpacer()))
	//_ = middle
	//l := widget.NewLabel("middle := container.New(layout.NewCenterLayout(), p.Content(c.win, &c.model))")
	//l.Wrapping = fyne.TextWrapWord

	//upper := container.NewHBox(left, widget.NewSeparator(), middle)
	//left := container.NewHBox(left, widget.NewSeparator(), l)
	upper := container.NewBorder(nil, nil, container.NewHBox(left, widget.NewSeparator()), nil, middle)
	quitButton := widget.NewButton("❌ Quit", c.Quit)
	prevButton := widget.NewButton("< Back", c.Prev)
	if c.current == 0 {
		prevButton.Disable()
	}
	nextButton := widget.NewButton("Next >", c.Next)
	if c.current == len(c.pages)-1 {
		nextButton.Disable()
	}

	buttons := container.NewBorder(nil, nil, quitButton,
		container.NewHBox(prevButton, nextButton))
	bottom := container.NewVBox(widget.NewSeparator(), buttons)
	_ = bottom

	//return container.NewBorder(nil, bottom, nil, nil, l)
	return container.NewBorder(nil, bottom, nil, nil, upper)
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
	}
	c.win = c.app.NewWindow("Network Security Inspect/Bypass Switch")
	c.win.Resize(fyne.NewSize(600, 400))
	c.win.SetFixedSize(true)
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
