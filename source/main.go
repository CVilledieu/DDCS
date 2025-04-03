package main

import (
	fyneApp "fyne.io/fyne/v2/app"
	fyneWidget "fyne.io/fyne/v2/widget"
)

func main() {
	app := fyneApp.New()
	win := app.NewWindow("Character Sheet")

	can := win.Canvas()
	can.SetContent(fyneWidget.NewLabel("Character Select"))

	win.ShowAndRun()
}
