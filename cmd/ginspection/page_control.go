package main

import (
	"context"
	"inspection/pkg/cone"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PageControl struct {
	statusLabel   *widget.Label
	bypassButton  *widget.Button
	inspectButton *widget.Button
}

var _ Page = &PageOptions{}

func (p *PageControl) Name() string {
	return "control"
}

func (p *PageControl) GetStatus(model *Model) {
	//	infinite := widget.NewProgressBarInfinite()
	cOne := model.COne()
	for {
		stop := InfiniteProgress(p.statusLabel)
		status, err := cOne.GetInspectionBypassStatus(context.TODO())
		stop()
		if err != nil {
			log.Println("err: ", err)
			p.statusLabel.SetText(err.Error())
			p.bypassButton.Enable()
			p.inspectButton.Enable()
			return
		} else {
			log.Println("STATUS", status)
			p.statusLabel.SetText(status.Status)
			log.Println("set status ", status.Status)
			switch status.Status {
			case "in-progress":
				time.Sleep(1 * time.Second)
				continue
			case "bypass":
				p.bypassButton.Disable()
				p.inspectButton.Enable()
				return
			case "inspect":
				p.bypassButton.Enable()
				p.inspectButton.Disable()
				return
			default:
				p.statusLabel.SetText(status.Error)
				return
			}
		}
	}
}

func (p *PageControl) Content(win fyne.Window, model *Model) fyne.CanvasObject {
	p.statusLabel = widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	go p.GetStatus(model)
	bypassFunc := func(action cone.Action) {
		cOne := model.COne()
		stop := InfiniteProgress(p.statusLabel)
		err := cOne.SetInspectionBypass(context.TODO(), action)
		stop()
		if err != nil {
			log.Println(err)
			return
		}
		p.GetStatus(model)
	}
	p.bypassButton = widget.NewButton("Bypass", func() {
		bypassFunc(cone.ActionBypass)
	})
	p.bypassButton.Disable()
	p.inspectButton = widget.NewButton("Inspect", func() {
		bypassFunc(cone.ActionInspect)
	})
	p.inspectButton.Disable()
	return widget.NewForm(
		widget.NewFormItem("Current State:", p.statusLabel),
		widget.NewFormItem("Bypass:", p.bypassButton),
		widget.NewFormItem("Inspect:", p.inspectButton),
	)
}

func (p *PageControl) AquireData(model *Model) error {
	return nil
}
