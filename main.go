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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 1280
var h size = 720

func makeEnv() {
	// black := color.Black
	paragraphEntry := widget.NewMultiLineEntry()
	const (
		loremIpsum = `Eu exercitation pariatur labore deserunt nulla dolore dolore quis sit ad deserunt dolore. Cillum anim non minim excepteur voluptate ex. Id mollit adipisicing officia dolore minim officia culpa aliquip. Ex ipsum ex dolor deserunt ut nisi velit incididunt voluptate sit duis.`
	)
	// newParagraph := canvas.NewText("", black)
	paragraphEntry.SetText(loremIpsum)
	paragraphEntry.Wrapping = fyne.TextTruncate
	// win.SetMainMenu(menu)
	// validateParag := widget.NewButton("validate paragraph", func() {
	// 	// entry, _ := paragraphEntry.
	// })
	content := container.NewVBox(
		paragraphEntry,
		// validateParag,
	)
	win.SetContent(content)
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
