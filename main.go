package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

// "fyne.io/fyne/container"

// "log"

// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/theme"

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 200
var h size = 200

func makeEnv() {
	r, _ := fyne.LoadResourceFromURLString("https://avatars.githubusercontent.com/u/36045855?s=200&v=4")
	// ceFromURLString("")
	// win.SetIcon(r)
	win.SetContent(canvas.NewImageFromResource(r))
}

type size = float32

func main() {
	launchApp()

}

func launchApp() {

	win.SetIcon(theme.FyneLogo())
	// a.Settings().SetTheme(theme.LightTheme())
	win.Resize(fyne.NewSize(w, h))

	makeEnv()
	win.ShowAndRun()
}
