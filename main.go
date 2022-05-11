package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 1280
var h size = 720

func makeEnv() {
	// black := color.Black
	// sub items
	// new := fyne.NewMenuItem("New", func() { fmt.Println("New file") })
	// save := fyne.NewMenuItem("Save", func() { fmt.Println("Save file") })
	// open := fyne.NewMenuItem("Open", nil)
	// items
	separator := fyne.NewMenuItemSeparator()
	file := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("New file") }),
		fyne.NewMenuItem("Save", func() { fmt.Println("Save file") }),
		fyne.NewMenuItem("Open", nil),
		separator,
		fyne.NewMenuItem("info", func() { fmt.Println("About GO/FYNE") }),
	)
	edit := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Undo", func() { fmt.Println("Undo") }),
		fyne.NewMenuItem("Redo", func() { fmt.Println("Redo") }),
	)

	// {
	// 	Label: "File",
	// 	Items: nil,
	// }
	// menu
	menu := fyne.NewMainMenu(
		file,
		edit,
		// today,
	)

	win.SetMainMenu(menu)

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
