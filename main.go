package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	colors "fyne-test/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
	redRect := canvas.NewRectangle(colors.Red)
	redRect.SetMinSize(stdSize)
	// redRect.Text = "red box"
	blueRect := canvas.NewRectangle(colors.Blue)
	blueRect.SetMinSize(stdSize)

	// scroller
	disabledBTN := widget.NewButton("Disabled", nil)
	// text := canvas.NewText("TEXT TO DISPLAY", color.Black)
	disabledBTN.Disable()
	box := container.NewHBox(
		blueRect,
		redRect,
	)

	scroll := container.NewVScroll(
		box,
	)

	// scroll direction
	scroll.Direction = container.ScrollHorizontalOnly

	// content := container.NewVBox(disabledBTN)
	win.SetContent(scroll)
}

type size = float32

func main() {
	launchApp()

}

func launchApp() {

	win.SetIcon(theme.FyneLogo())
	// a.Settings().SetTheme(theme.LightTheme())
	win.Resize(stdSize)
	win.CenterOnScreen()

	makeEnv()
	win.ShowAndRun()
}
