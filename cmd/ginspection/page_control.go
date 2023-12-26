package main

import (
	"context"
	"inspection/pkg/cone"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type PageControl struct {
	statusLabel   *widget.Label
	bypassButton  *widget.Button
	inspectButton *widget.Button
}

var _ Page = &PageOptions{}

func (p *PageControl) Name() string {
	return "Control"
}

func (p *PageControl) GetStatus(model *Model) {
	cOne := model.COne()
	for {
		stop := InfiniteProgress(p.statusLabel)
		status, err := cOne.GetInspectionBypassStatus(context.TODO(), model.config.AWSRegion)
		stop()
		if err != nil {
			p.statusLabel.SetText(err.Error())
			p.bypassButton.Enable()
			p.inspectButton.Enable()
			return
		} else {
			switch status.Status {
			case cone.StatusIn_progress:
				time.Sleep(1 * time.Second)
				continue
			case cone.StatusFail:
				p.bypassButton.Enable()
				p.inspectButton.Enable()
				p.statusLabel.SetText(cone.StatusFail.String() + ": " + status.Error)
				return
			case cone.StatusSuccess:
				switch status.Action {
				case cone.ActionBypass:
					p.bypassButton.Disable()
					p.inspectButton.Enable()
				case cone.ActionInspect:
					p.bypassButton.Enable()
					p.inspectButton.Disable()
				}
				p.statusLabel.SetText(status.Action.String())
				return
			default:
				p.statusLabel.SetText(status.Error)
				p.bypassButton.Enable()
				p.inspectButton.Enable()
				return
			}
		}
	}
}

func (p *PageControl) Content(win fyne.Window, model *Model) fyne.CanvasObject {
	p.statusLabel = widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	go p.GetStatus(model)
	controlFunc := func(action cone.Action) {
		cOne := model.COne()
		stop := InfiniteProgress(p.statusLabel)
		err := cOne.SetInspectionBypass_(context.TODO(), model.config.AWSRegion, action)
		stop()
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		p.GetStatus(model)
	}
	p.bypassButton = widget.NewButton("Bypass", func() {
		controlFunc(cone.ActionBypass)
	})
	p.inspectButton = widget.NewButton("Inspect", func() {
		controlFunc(cone.ActionInspect)
	})
	return container.NewVBox(
		container.NewHBox(widget.NewLabel("Current State:"), p.statusLabel),
		p.bypassButton,
		p.inspectButton,
	)
}

func (p *PageControl) AquireData(model *Model) error {
	return nil
}
