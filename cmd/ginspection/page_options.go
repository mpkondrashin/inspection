package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type PageOptions struct {
	//regionList     *widget.Select
	accountIDEntry *widget.Entry
	apiKeyEntry    *widget.Entry
	awsRegionList  *widget.Select
}

var _ Page = &PageOptions{}

func (p *PageOptions) Name() string {
	return "options"
}

func (p *PageOptions) Content(win fyne.Window, model *Model) fyne.CanvasObject {
	labelTop := widget.NewLabel("Provide all nessesary options")
	/*
		detectButton := widget.NewButton("detect", nil)
		detectButton.OnTapped = func() {
			detectButton.Disable()
			saveText := detectButton.Text
			stop := InfiniteProgressFunc(func(s string) {
				detectButton.SetText(s)
			})
			cOne := model.COne()
			info, err := cOne.GetAccountInfo(context.TODO())
			stop()
			detectButton.SetText(saveText)
			detectButton.Enable()
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			p.regionList.SetSelected(info.Region)
		}
	*/
	p.accountIDEntry = widget.NewEntry()
	p.accountIDEntry.Text = model.config.AccountID
	p.accountIDEntry.Validator = ValidateAccountID
	p.accountIDEntry.PlaceHolder = "Go to Administration-Account Settings->ID"
	//accountIDOk := false
	//apiKeyOk := false

	/*p.accountIDEntry.OnChanged = func(s string) {
		accountIDOk = p.accountIDEntry.Validate() == nil
		if accountIDOk && apiKeyOk {
			detectButton.Enable()
		} else {
			detectButton.Disable()
		}
	}*/

	p.apiKeyEntry = widget.NewPasswordEntry()
	p.apiKeyEntry.Text = model.config.apiKeyDecrypted
	p.apiKeyEntry.PlaceHolder = "Go to Administration->API Keys->New"
	/*p.accountIDEntry.OnChanged = func(s string) {
		apiKeyOk = s != ""
		if accountIDOk && apiKeyOk {
			detectButton.Enable()
		} else {
			detectButton.Disable()
		}
	}*/
	/*
		p.regionList = widget.NewSelect(CloudOneRegions, nil)
		p.regionList.SetSelected(model.config.Region)

		regionHBox := container.NewHBox(p.regionList, detectButton)
	*/
	/*
		p.awsRegionList = widget.NewSelect(NSHIRegions, nil)
		p.awsRegionList.SetSelected(model.config.AWSRegion)
		awsRegionsHBox := container.NewHBox(p.awsRegionList, widget.NewButton("detect", nil))
	*/
	optionsForm := widget.NewForm(
		widget.NewFormItem("AccoutID:", p.accountIDEntry),
		widget.NewFormItem("API Key:", p.apiKeyEntry),
		//widget.NewFormItem("Cloud One Region:", regionHBox),
		//widget.NewFormItem("AWS NSHI Region:", awsRegionsHBox),
	)
	return container.NewVBox(labelTop, optionsForm)
}

func (p *PageOptions) AquireData(model *Model) error {
	model.config.apiKeyDecrypted = p.apiKeyEntry.Text
	model.config.AccountID = p.accountIDEntry.Text

	cOne := model.COne()
	info, err := cOne.GetAccountInfo(context.TODO())

	if err != nil {
		return err // dialog.ShowError(err, win)
	}
	//	p.regionList.SetSelected(info.Region)

	model.config.Region = info.Region // p.regionList.Selected
	//model.config.AWSRegion = p.awsRegionList.Selected
	if model.Changed() {
		return model.Save(configFileName)
	}
	return nil
}
