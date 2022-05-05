package main

import (
	"fmt"
	colors "fyne-test/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 800
var h size = 450

func makeEnv() {
	var mainDisplay *fyne.Container
	// var secDisplay *fyne.Container
	fmt.Printf("%T\n", mainDisplay)
	hori := true
	heading := canvas.NewText("This is my title", colors.Black)
	parag := canvas.NewText("This is my paragraph", colors.Black)
	line := canvas.NewLine(colors.Gray)
	line.StrokeWidth = 2
	horiBox := container.NewHBox(
		heading,
		line,
		parag,
	)
	vertBox := container.NewVBox(
		heading,
		line,
		parag,
	)
	if hori {
		win.SetContent(horiBox) // will display from top left to bottom right

	} else {
		win.SetContent(vertBox) // will display from top left to bottom right

	}

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
