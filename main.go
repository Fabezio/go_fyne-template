package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	black := color.Black
	label1 := canvas.NewText("Title 1", black)
	label2 := canvas.NewText("Title 2", black)
	win.SetContent(container.NewHSplit(
		label1,
		label2,
	))
}

type size = float32

func main() {
	launchApp()

}

func launchApp() {
	a.Settings().SetTheme(theme.LightTheme())
	win.Resize(fyne.NewSize(w, h))

	makeEnv()
	win.ShowAndRun()
}
