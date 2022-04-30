package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("window 1")
	const w float32 = 1280
	const h float32 = 800
	win.SetContent(widget.NewLabel("Hello World!\n\nGood to see you"))
	win.Resize(fyne.NewSize(w, h))
	win.SetMaster()
	win.Show()

	win2 := a.NewWindow("window 2")
	win2.SetContent(widget.NewButton("Click here", func() {
		win3 := a.NewWindow("dependency window")
		win3.SetContent(widget.NewLabel("What do you want to place here ?"))
		win3.Resize(fyne.NewSize(150, 100))
		win3.Show()
	}))
	win2.Resize(fyne.NewSize(w/2, h/2))
	win2.Show()

	a.Run()
}
