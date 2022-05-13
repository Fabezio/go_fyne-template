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
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 1360
var h size = 720

var stdSize = fyne.NewSize(w, h)

func makeEnv() {
	// boxes

	win.SetContent(container.NewVBox(
		coloredButton("open .txt files", colors.Green),
	))
}

type size = float32

func main() {
	// colors.Colors()
	launchApp()

}

func coloredButton(text string, colorName color.NRGBA) *fyne.Container {
	btn := widget.NewButton(text, func() {
		fileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(r)
				result := fyne.NewStaticResource("name", data)
				entry := widget.NewMultiLineEntry()
				entry.SetText(string(result.StaticContent))
				win := fyne.CurrentApp().NewWindow(
					string(result.StaticName))
				win.Resize(stdSize)
				win.SetContent(container.NewScroll(entry))
				win.Show()

			}, win)
		fileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		fileDialog.Show()
		fileDialog.Resize(stdSize)
	})

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
	win.Show()
	a.Run()
}
