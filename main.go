package main

import (
	colors "fyne-test/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func makeEnv() {
	horizontalGradient := canvas.NewHorizontalGradient(colors.Orange, colors.Yellow)
	verticalGradient := canvas.NewVerticalGradient(colors.Orange, colors.Yellow)
	linearGradient := canvas.NewLinearGradient(colors.Orange, colors.Black, 45)
	radialGradient := canvas.NewRadialGradient(colors.Black, colors.Orange)
	w1.SetContent(horizontalGradient) // will display from top left to bottom right
	w1.Show()
	w2.SetContent(verticalGradient) // will display from top left to bottom right
	w2.Show()
	w3.SetContent(linearGradient) // will display from top left to bottom right
	w3.Show()
	w4.SetContent(radialGradient) // will display from top left to bottom right
	w4.Show()

}

type size = float32

var title = "a mere http url"
var a = app.New()
var w1 = a.NewWindow("Horizontal gradient")
var w2 = a.NewWindow("Vertical gradient")
var w3 = a.NewWindow("Linear gradient")
var w4 = a.NewWindow("Radial gradient")
var w size = 200
var h size = 200

func main() {
	launchApp()

}

func launchApp() {
	w1.Resize(fyne.NewSize(w, h))
	w2.Resize(fyne.NewSize(w, h))
	w3.Resize(fyne.NewSize(w, h))
	w4.Resize(fyne.NewSize(w, h))
	makeEnv()
	a.Run()
}
