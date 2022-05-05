package main

import (
	colors "fyne-test/helpers"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeEnv() {
	black := color.Black
	gray := colors.Gray
	heading := canvas.NewText("visit my website", black)
	heading.TextSize = 30
	line := canvas.NewLine(gray)

	// parag := canvas.NewText("select your gender", black)
	url, _ := url.Parse("https://github.com/Fabezio")
	hyperlink := widget.NewHyperlink("my github repos", url)

	// make vertical layout (VBox)
	layout := container.NewVBox(

		heading,
		line,
		hyperlink,
	)
	win.SetContent(layout)

}

type size = float32

var title = "a mere http url"
var a = app.New()
var win = a.NewWindow(title)
var w size = 400
var h size = 400

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))
	makeEnv()
	win.ShowAndRun()
}
