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
	// heading := canvas.NewText("Random Nb Gen", colors.Black)
	// heading.TextSize = 40
	face1 := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-one-solid.svg"
	face2 := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-two-solid.svg"
	face3 := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-three-solid.svg"
	face4 := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-four-solid.svg"
	face5 := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-five-solid.svg"
	face6 := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-six-solid.svg"
	dice := "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice/dice-d6-solid.svg"

	faces := [6]string{face1, face2, face3, face4, face5, face6}

	diceImg := canvas.NewImageFromFile(dice)
	diceImg.FillMode = canvas.ImageFillOriginal

	// img := canvas.NewImageFromFile("/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/dice-four-solid.svg")
	btn := widget.NewButton("play", func() {
		fmt.Println("generating...")

		time.Sleep(250 * time.Millisecond)
		rand := rand.Intn(6)
		fmt.Println(rand + 1)
		diceImg.File = faces[rand]
		diceImg.FillMode = canvas.ImageFillOriginal
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
