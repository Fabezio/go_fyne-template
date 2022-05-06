package main

import (

	// "fyne.io/fyne/container"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/theme"
)

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 400
var h size = 400

func myIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Printf("%T\n", err)
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {

		fmt.Printf("%T\n", err)
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Query
}

func myCity() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Printf("%T\n", err)
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {

		fmt.Printf("%T\n", err)
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)
	return ip.City
}
func myCountry() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Printf("%T\n", err)
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {

		fmt.Printf("%T\n", err)
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)
	return ip.Country
}

type IP struct {
	Query   string
	Country string
	City    string
}

func makeEnv() {
	whichIP := widget.NewLabel("What is my IP ?")
	labelIP := widget.NewLabel("Your IP is ...")
	valueCountry := widget.NewLabel("...")
	values := widget.NewCard(
		"...",
		"...",
		valueCountry,
	)
	btn := widget.NewButton("Search IP", func() {
		valueCountry.Text = myCountry()
		valueCountry.Refresh()
		values.Title = myIP()
		values.Subtitle = myCity()
		values.Content = valueCountry
		values.Refresh()
	})

	win.SetContent(container.NewVBox(
		whichIP,
		labelIP,
		values,
		btn,
	))

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
