package main

import (
	"fmt"
	colors "fyne-test/helpers"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 800
var h size = 450

func makeEnv() {
	heading := canvas.NewText("Random Nb Gen", colors.Black)
	heading.TextSize = 40
	btn := widget.NewButton("click here to generate a number", func() {
		fmt.Println("generating...")

		time.Sleep(250 * time.Millisecond)
		rand := rand.Intn(100)
		fmt.Println(rand)
		heading.Text = fmt.Sprintf("number is %v", rand)
		heading.Color = colors.Orange
		heading.Refresh()
	})

	cont := container.NewVBox(btn, heading)
	win.SetContent(cont)

}

type size = float32

func main() {
	launchApp()

}

func launchApp() {
	win.Resize(fyne.NewSize(w, h))

	makeEnv()
	win.ShowAndRun()
}
