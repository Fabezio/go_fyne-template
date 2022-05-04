package main

import (
	"fmt"
	colors "fyne-test/helpers"

	// colors "fyne-test/helpers"
	// "image/color"

	// "fyne.io/fyne/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/canvas"
)

type size = float32

func main() {
	title := "Gradient: tutorial"
	fmt.Println("lauching...")
	var a = app.New()
	win := a.NewWindow(title)
	// const w size = 600
	// const h size = 450

	const maxW size = 1440
	const maxH size = 900

	// var assets = "/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/"
	// var setColor = color.NRGBA{R: 127, G: 0, B: 255, A: 200}
	// var text = canvas.NewText("it's good to see y'all !", setColor)

	// var checkbox = widget.NewCheck("my checkbox", func(b bool) {
	// 	str := "bad"
	// 	if b {
	// 		str = "good"
	// 	}

	// 	fmt.Println(str)
	// })

	// var image = canvas.NewImageFromFile(assets + "NicePng_languages-icon-png_3801538.png")

	// var circle = canvas.NewCircle(color.NRGBA{R: red, G: green, B: blue, A: alpha})
	// var rect = canvas.NewRectangle(colors.White)
	// rect.StrokeWidth = 3
	// rect.StrokeColor = colors.Cyan

	// line := canvas.NewLine(colors.Blue)
	// line.StrokeWidth = 3
	// line.StrokeColor = colors.Amethyst

	// icon := widget.NewIcon(theme.ColorPaletteIcon())

	// gradient := canvas.NewRadialGradient(colors.Yellow, color.Black)
	btn := widget.NewButton("click me", func() {
		fmt.Println("OK")
	})
	label := widget.NewLabel("Hello world !")
	// label.StrokeColor = colors.Red
	// label.TextStyle =
	parag := canvas.NewText("Glad to see you.", colors.Black)
	parag2 := canvas.NewText("Hoping you too.", colors.Blue)

	vBox := container.NewVBox(btn,
		label, parag, parag2)
	// hBox := container.NewHBox(btn,
	// label)

	win.SetContent(vBox)

	// win.Content().Position(fyne.Position(100, 100))

	win.Resize(fyne.NewSize(maxH*3/4, maxH*3/4))

	win.ShowAndRun()

	// a.Run()
}
