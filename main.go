package main

import (

	// "fyne.io/fyne/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 400
var h size = 400

func makeEnv() {

	img := canvas.NewImageFromFile("/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-4-solid.svg")
	img.FillMode = canvas.ImageFillOriginal
	card := widget.NewCard(
		"title",
		"subtitle",
		img,
	)

	win.SetContent(card)

}

type size = float32

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))

	makeEnv()
	win.ShowAndRun()
}
