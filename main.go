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
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 500
var h size = 720

func makeEnv() {
	// Let create items for accoridan
	// first argument is title, 2nd is description/details
	item1 := widget.NewAccordionItem("A",
		widget.NewLabel("A for Apple"))
	item2 := widget.NewAccordionItem("B",
		widget.NewLabel("B for Ball"))
	item3 := widget.NewAccordionItem("C",
		widget.NewLabel("C for Cat"))
	ac := widget.NewAccordion(item1, item2, item3)
	win.SetContent(ac)
}

type size = float32

func main() {
	launchApp()

}

func launchApp() {

	win.SetIcon(theme.FyneLogo())
	// a.Settings().SetTheme(theme.LightTheme())
	win.Resize(fyne.NewSize(w, h))
	win.CenterOnScreen()

	makeEnv()
	win.ShowAndRun()
}
