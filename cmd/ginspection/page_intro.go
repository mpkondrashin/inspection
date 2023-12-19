package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	Title = `GInspection v0.0.1`

	IntoText = `Version v0.0.1

Trend Micro Cloud One Netwrok Security Hosted Infrastructure controlling utility.

GInspection allows to turn on and off bypass mode for  Netwrok Security Hosted Infrastructure in given AWS region`

	License = `MIT License

Copyright (c) 2024 Michael Kondrashin (mkondrashin@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`
)

type PageIntro struct {
}

var _ Page = &PageIntro{}

func (p *PageIntro) Name() string {
	return "info"
}

func (p *PageIntro) Content(win fyne.Window, model *Model) fyne.CanvasObject {

	title := widget.NewLabelWithStyle("GInspection", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	report := widget.NewRichTextFromMarkdown(IntoText)
	report.Wrapping = fyne.TextWrapWord
	link, _ := url.Parse("https://github.com/mpkondrashin/inspection")
	github := widget.NewHyperlink("GInspector repository on GitHub", link)

	return container.NewVBox(
		title,
		report,
		github,
		widget.NewButton("License Info", func() {
			licenseLabel := widget.NewLabel(License)
			sc := container.NewScroll(licenseLabel)
			popup := dialog.NewCustom("License Information", "Close", sc, win) //fyne.CurrentApp().Driver().Canvas())
			popup.Resize(fyne.NewSize(800, 600))
			popup.Show()
		}),
	)
}

func (p *PageIntro) AquireData(model *Model) error {
	return nil
}
