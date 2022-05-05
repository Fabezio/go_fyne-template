package main

import (
	colors "fyne-test/helpers"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeEnv() {
	black := color.Black
	gray := colors.Gray
	parag := canvas.NewText("select your gender", black)
	checkbox := widget.NewCheck("you're a male", func(b bool) {

		parag.Refresh()
		if b {
			parag.Text = "exactly"
		} else {
			parag.Text = "No, i'm a girl"

		}
		parag.Refresh()
	})
	heading := canvas.NewText("Welcome", black)
	heading.TextSize = 30

	line := canvas.NewLine(gray)
	// make vertical layout (VBox)
	layout := container.NewVBox(
		checkbox,
		heading,
		line,
		parag,
	)
	win.SetContent(layout)

}

type size = float32

var title = "First fyne app"
var a = app.New()
var win = a.NewWindow(title)
var w size = 400
var h size = 400

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))
	makeEnv()
	win.ShowAndRun()
}
