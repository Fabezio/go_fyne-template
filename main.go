package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 400
var h size = 400

func makeEnv() {
	// var height float64
	// var weight float64

	// hStr := ""
	// wStr := ""

	msg := canvas.NewText("", color.Black)
	// fmt.Println(bmi)
	// input := widget.NewForm(fyne.)
	hInput := widget.NewEntry()
	wInput := widget.NewEntry()
	hInput.SetPlaceHolder("Height?")
	wInput.SetPlaceHolder("Weight?")
	bmiCalc := widget.NewButton("Calculate BMI", func() {
		height, _ := strconv.ParseFloat(hInput.Text, 64)
		weight, _ := strconv.ParseFloat(wInput.Text, 64)
		var bmi float64 = weight / (math.Pow(height/100, 2))
		msg.Text = fmt.Sprintf("%v kg/m2.  ", bmi)
		msg.Refresh()
		if bmi <= 24.9 {
			msg.Text += "U're healthy"
		} else if bmi <= 29.9 {
			msg.Text += "U're over weight"
		} else if bmi <= 34.9 {
			msg.Text += "U're severely over weight"
		} else if bmi <= 39.9 {
			msg.Text += "U're obese"
		} else {
			msg.Text += "U're severely obese"

		}
		msg.Refresh()
		// fmt.Print("you're done!")
	})
	win.SetContent(container.NewVBox(
		hInput,
		wInput,
		bmiCalc,
		msg,
	))
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
