package main

import (
	colors "fyne-test/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func makeEnv() {
	hori := true
	line := canvas.NewLine(colors.Gray)

	line.StrokeWidth = 2
	layout := container.NewVBox(line)
	// choose if you want to display a single line or in a box (hozizontal line)
	if hori {
		win.SetContent(layout) // will display a separation line
	} else {
		win.SetContent(line) // will display from top left to bottom right
	}

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
