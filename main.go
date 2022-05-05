package main

import (
	colors "fyne-test/helpers"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func makeEnv() {
	circle := canvas.NewCircle(colors.Orange)
	circle.StrokeColor = color.Black
	circle.StrokeWidth = 2

	win.SetContent(circle)

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
