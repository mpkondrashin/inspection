package main

import (
	"fyne.io/fyne/v2"
)

// https://cloudone.trendmicro.com/docs/identity-and-account-management/c1-regions/
/*
var CloudOneRegions = []string{
	"us-1",
	"in-1",
	"gb-1",
	"jp-1",
	"de-1",
	"au-1",
	"ca-1",
	"sg-1",
	"trend-us-1",
}
*/

const configFileName = "config.yaml"

type Page interface {
	Name() string
	Content(win fyne.Window, model *Model) fyne.CanvasObject
	AquireData(model *Model) error
}

/*
func InfiniteProgress(label *widget.Label) func() {
	stop := make(chan struct{})
	go func() {
		chars := ".oOo"
		chars = `/-\|`
		chars = ` booo dooo oboo odoo oobo oodo ooob oood oooq ooop ooqo oopo oqoo opoo qooo pooo`
		//substr := chars[i : i+4]
		i := 0
		for {
			select {
			case <-stop:
				//fmt.Println("got a close signal")
				//wg.Done()
				return
			default:
				label.SetText(chars[i*5 : i*5+5])
				//				label.SetText(string(chars[i]))
				i++
				if i*5 == len(chars) {
					i = 0
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return func() {
		stop <- struct{}{}
	}
}*/

func main() {
	c := NewNSHIControl()
	c.Run()
}
