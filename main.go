package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeEnv() {
	icon := widget.NewIcon(theme.AccountIcon())

	win.SetContent(icon) // will display from top left to bottom right

}

type size = float32

var title = "a mere http url"
var a = app.New()
var win = a.NewWindow(title)
var w size = 500
var h size = 500

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))
	makeEnv()
	win.ShowAndRun()
}
