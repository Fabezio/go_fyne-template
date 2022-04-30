package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var a = app.New()

func main() {
	win := a.NewWindow("window 1")
	const w float32 = 1280
	const h float32 = 800
	win.SetContent(widget.NewLabel("Hello World!\n\nGood to see you"))
	win.SetContent(widget.NewButton("Click here", func() {
		btn := a.NewWindow("dependency window")
		btn.SetContent(widget.NewLabel("Well done!"))
		btn.Resize(fyne.NewSize(150, 100))
		btn.Show()
	}))
	win.Resize(fyne.NewSize(w, h))
	// win.SetMaster()
	win.Show()

	a.Run()
}
