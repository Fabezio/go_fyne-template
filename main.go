package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func makeEnv() {
	img := "assets/tux.png"
	curPath := "/Users/fabriceriquet/Documents/dev/go/test-fyne/"
	file := canvas.NewImageFromFile(curPath + img)

	win.SetContent(file)

}

type size = float32

var title = "a mere http url"
var a = app.New()
var win = a.NewWindow(title)
var w size = 640
var h size = 768

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))
	makeEnv()
	win.ShowAndRun()
}
