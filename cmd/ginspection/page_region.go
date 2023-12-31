package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var NSHIRegions = []string{
	"us-west-1",
	"us-west-2",
	"us-east-1",
	"us-east-2",
	"ap-south-1",
	"sa-east-1",
	"ap-south-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-northeast-3",
	"eu-central-1",
	"eu-north-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"ca-central-1",
}

type PageRegion struct {
	//regionList     *widget.Select
	awsRegionList *widget.Select
}

var _ Page = &PageRegion{}

func (p *PageRegion) Name() string {
	return "Region"
}

func (p *PageRegion) Content(win fyne.Window, model *Model) fyne.CanvasObject {

	selectedRegion := ""
	if model.config.AWSRegion != "" {
		selectedRegion = model.config.AWSRegion
	} else {
		awsRegions := model.COne().DetectAWSRegions(context.TODO())
		if len(awsRegions) > 0 {
			selectedRegion = awsRegions[0]
		}
	}

	labelTop := widget.NewLabel("Choose AWS Region")

	p.awsRegionList = widget.NewSelect(NSHIRegions, nil)
	if selectedRegion != "" {
		p.awsRegionList.SetSelected(selectedRegion)
	}
	passwordForm := widget.NewForm(
		widget.NewFormItem("AWS NSHI Region:", p.awsRegionList),
	)
	return container.NewVBox(labelTop, passwordForm)
}

func (p *PageRegion) AquireData(model *Model) error {
	model.config.AWSRegion = p.awsRegionList.Selected
	if model.Changed() {
		return model.Save()
	}
	return nil
}
