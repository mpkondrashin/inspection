package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type PagePassword struct {
	passwordEntry *widget.Entry
}

var _ Page = &PagePassword{}

func (p *PagePassword) Name() string {
	return "auth"
}

func (p *PagePassword) Content(win fyne.Window, model *Model) fyne.CanvasObject {
	labelTop := widget.NewLabel("Provide password that will be used to encrypt/decrypt API key")
	p.passwordEntry = widget.NewPasswordEntry()
	p.passwordEntry.Text = model.password
	p.passwordEntry.Validator = CheckPassword
	passwordFormItem := widget.NewFormItem("Password:", p.passwordEntry)
	passwordForm := widget.NewForm(passwordFormItem)
	return container.NewVBox(labelTop, passwordForm)
}

func (p *PagePassword) AquireData(model *Model) error {
	//if len(p.passwordEntry.Text) < 1 {
	//	return errors.New("password is too short")
	//}
	model.password = p.passwordEntry.Text
	err := model.config.Load(configFileName, model.password)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
