package main

import (
	"fmt"
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

var w size = 400
var h size = 400

func makeEnv() {
	dice := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-d6-solid.svg"

	diceImg := canvas.NewImageFromFile(dice)
	diceImg.FillMode = canvas.ImageFillOriginal

	btn := widget.NewButton("play", func() {
		fmt.Println("generating...")

		time.Sleep(250 * time.Millisecond)
		rand := rand.Intn(6)
		// str :=
		fmt.Println(rand + 1)
		diceImg.File = fmt.Sprintf("/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-%v-solid.svg", rand+1)

		diceImg.Refresh()
	})

	cont := container.NewVBox(
		diceImg,
		btn,
	)
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
