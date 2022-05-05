package main

import (
	"fmt"
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

	parag := canvas.NewText("This is a paragraph", black)
	btn := widget.NewButton("click me", func() {
		fmt.Println("yea!")
		parag.Text = "Well done!"
		parag.Color = colors.Teal
		parag.Refresh()
	})
	heading := canvas.NewText("Welcome", black)
	heading.TextSize = 30
	// label := widget.NewLabel("Hello World !")
	line := canvas.NewLine(gray)
	// make vertical layout (VBox)
	layout := container.NewVBox(
		btn,
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
