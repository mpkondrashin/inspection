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
	//awsRegionList  *widget.Select
}

var _ Page = &PageOptions{}

func (p *PageOptions) Name() string {
	return "Options"
}

func (p *PageOptions) Content(win fyne.Window, model *Model) fyne.CanvasObject {
	labelTop := widget.NewLabel("Provide all nessesary options")

	p.accountIDEntry = widget.NewEntry()
	p.accountIDEntry.Text = model.config.AccountID
	p.accountIDEntry.Validator = ValidateAccountID
	//p.accountIDEntry.PlaceHolder = "Go to Administration-Account Settings->ID"
	accountIDFormItem := widget.NewFormItem("AccoutID:", p.accountIDEntry)
	accountIDFormItem.HintText = "Go to Administration-Account Settings->ID"
	p.apiKeyEntry = widget.NewPasswordEntry()
	p.apiKeyEntry.Text = model.config.apiKeyDecrypted
	p.apiKeyEntry.PlaceHolder = "" // "Go to Administration->API Keys->New"
	apiKeyFormItem := widget.NewFormItem("API Key:", p.apiKeyEntry)
	apiKeyFormItem.HintText = "Go to Administration->API Keys->New"

	optionsForm := widget.NewForm(
		accountIDFormItem,
		apiKeyFormItem,
	)
	return container.NewVBox(labelTop, optionsForm)
}

func (p *PageOptions) AquireData(model *Model) error {
	model.config.apiKeyDecrypted = p.apiKeyEntry.Text
	model.config.AccountID = p.accountIDEntry.Text

	cOne := model.COne()
	info, err := cOne.GetAccountInfo(context.TODO())

	if err != nil {
		return err
	}

	model.config.Region = info.Region
	if model.Changed() {
		return model.Save(configFileName)
	}
	return nil
}
