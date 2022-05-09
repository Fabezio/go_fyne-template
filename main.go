package main

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

// package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// "fyne.io/fyne/container"

// "log"

// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/theme"

var title = "fyne tutorial"
var a = app.New()
var win = a.NewWindow(title)

var w size = 500
var h size = 500

func makeEnv() {
	lat := 48.856614
	lon := 2.3522219
	lang := "en"
	units := "metric"
	apiKey := "f6f03d4b430d972622542e12b63f852d"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&lang=%v&units=%v&appid=%v", lat, lon, lang, units, apiKey)
	// apikey :=
	res, err := http.Get(url)
	// https: //api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={API key}
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	forecast, err := UnmarshalWelcome(body)

	if err != nil {
		fmt.Println(err)
	}

	black := color.Black
	mainTitle := fmt.Sprintf("Current forecast in %v", forecast.Name)
	text := canvas.NewText(mainTitle, black)
	text.TextSize = 30
	tempStr := fmt.Sprintf("Temp: %v °C", forecast.Main.Temp)
	feltStr := fmt.Sprintf("Feels like: %v °C", forecast.Main.FeelsLike)
	maxStr := fmt.Sprintf("Max temp: %v °C", forecast.Main.TempMax)
	minStr := fmt.Sprintf("min temp: %v °C", forecast.Main.TempMin)
	tempLabel := canvas.NewText("Temperatures", black)
	tempLabel.TextSize = 30

	temp := canvas.NewText(tempStr, black)
	felt := canvas.NewText(feltStr, black)
	mintemp := canvas.NewText(minStr, black)
	maxtemp := canvas.NewText(maxStr, black)
	weather := canvas.NewText(forecast.Weather[0].Description, black)
	country := canvas.NewText(forecast.Sys.Country, black)
	iconURI := fmt.Sprintf("/Users/fabriceriquet/Documents/dev/go/test-fyne/assets/forecast/%v@2x.png", forecast.Weather[0].Icon)
	// fmt.Println(iconURI)
	icon := canvas.NewImageFromFile(iconURI)
	icon.FillMode = canvas.ImageFillOriginal

	win.SetContent(container.NewVBox(
		icon,
		text,
		weather,
		country,
		tempLabel,
		temp,
		felt,
		maxtemp,
		mintemp,
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

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
