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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 500
var h size = 720

var stdSize = fyne.NewSize(w, h)

func makeEnv() {
	// boxes

	win.SetContent(container.NewVBox(
		coloredButton("About", colors.Blue),
		coloredButton("Submit", colors.Orange),
		coloredButton("Reset", colors.Red),
		coloredButton("Reach",
			color.NRGBA{R: 0, G: 255, B: 255, A: 127}),
	))
}

type size = float32

func main() {
	colors.Colors()
	launchApp()

}

func coloredButton(text string, colorName color.NRGBA) *fyne.Container {
	btn := widget.NewButton(text, nil)

	btnColor := canvas.NewRectangle(colorName)

	container := container.New(
		layout.NewMaxLayout(),
		btnColor,
		btn,
	)

	return container
}

func launchApp() {

	win.SetIcon(theme.FyneLogo())
	// a.Settings().SetTheme(theme.LightTheme())
	win.Resize(stdSize)
	win.CenterOnScreen()

	makeEnv()
	win.ShowAndRun()
}
