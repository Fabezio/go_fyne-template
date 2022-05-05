package main

import (
	colors "fyne-test/helpers"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func makeEnv() {
	rect := canvas.NewRectangle(colors.Orange)
	rect.StrokeColor = color.Black
	rect.StrokeWidth = 2
	win.SetContent(rect)

}

type size = float32

var title = "a mere http url"
var a = app.New()
var win = a.NewWindow(title)
var w size = 640
var h size = 768

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))
	makeEnv()
	win.ShowAndRun()
}
