package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// "fyne.io/fyne/container"

// "log"

// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/theme"

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 200
var h size = 200

func makeEnv() {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "1234567890"
	speChars := "$£@#&§!?"
	passwdLength := 8
	genPasswdStr := ""
	genPasswd := canvas.NewText(genPasswdStr, color.Black)
	// ceFromURLString("")
	generator := widget.NewButton("generate Password", func() {
		genPasswdStr = ""

		// genPasswd.Text = ""
		// genPasswd.Refresh()
		for n := 0; n < passwdLength; n++ {
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(4)
			// fmt.Println(randNum)
			switch randNum {
			case 0:
				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(lower))
				genPasswdStr += string(lower[randIndex])
			case 1:
				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(upper))
				genPasswdStr += string(upper[randIndex])
			case 2:
				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(numbers))
				genPasswdStr += string(numbers[randIndex])
			case 3:
				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(speChars))
				genPasswdStr += string(speChars[randIndex])

			}
		}
		genPasswd.Text = genPasswdStr
		genPasswd.Refresh()
	})

	// fmt.Println(genPasswdStr)

	// win.SetIcon(r)
	win.SetContent(container.NewVBox(
		generator,
		genPasswd,
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
