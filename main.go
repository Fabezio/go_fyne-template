package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	colors "fyne-test/helpers"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// "fyne.io/fyne/container"

// "log"

// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/theme"

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 400
var h size = 400

func makeEnv() {
	randColor := color.Black
	// randRed := rand.Intn(255)
	// white := color.White
	line := canvas.NewLine(colors.Gray)
	line.StrokeWidth = 2
	rect := canvas.NewRectangle(randColor)
	rect.SetMinSize(fyne.NewSize(w-10, h/2))
	randomColorBTN := widget.NewButton("random color", func() {
		rect.FillColor = color.NRGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
		rect.Refresh()
	})
	randomRed := widget.NewButton("random red", func() {
		rect.FillColor = color.NRGBA{R: uint8(rand.Intn(255)), G: 0, B: 0, A: 255}
		rect.Refresh()
	})
	randomGreen := widget.NewButton("random green", func() {
		rect.FillColor = color.NRGBA{R: 0, G: uint8(rand.Intn(255)), B: 0, A: 255}
		rect.Refresh()
	})
	randomBlue := widget.NewButton("random blue", func() {
		rect.FillColor = color.NRGBA{R: 0, G: 0, B: uint8(rand.Intn(255)), A: 255}
		rect.Refresh()
	})

	// label1 := canvas.NewText("Title 1", black)
	// label2 := canvas.NewText("Title 2", white)
	win.SetContent(container.NewVBox(
		rect,
		// line,
		randomColorBTN,
		line,
		randomRed,
		randomGreen,
		randomBlue,
	))
}

type size = float32

func main() {
	launchApp()

}

func launchApp() {
	// a.Settings().SetTheme(theme.LightTheme())
	win.Resize(fyne.NewSize(w, h))

	makeEnv()
	win.ShowAndRun()
}
