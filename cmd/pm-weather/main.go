package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/resty.v1"
)

var (
	token  = kingpin.Flag("token", "API token for OpenWeatherMap").Short('t').Required().String()
	cityID = kingpin.Flag("city-id", "API token for OpenWeatherMap").Short('c').Required().String()
	spaces = kingpin.Flag("spaces", "Spaces between the icon and the text").Default("1").Int()
	debug  = kingpin.Flag("debug", "Print out errors").Default("false").Bool()
)

func main() {
	kingpin.Parse()

	if err := run(); err != nil {
		if *debug {
			fmt.Println(err)
		} else {
			fmt.Println("ERR")
		}
		os.Exit(1)
	}
}

func run() error {
	temperatureIcon := "%{F#d08770}%{F-}"
	rhIcon := "%{F#81a1c1}%{F-}"
	sunIcon := "%{F#ebcb8b}%{F-}"
	spacesString := strings.Repeat(" ", *spaces)
	doubleSpacesString := strings.Repeat(" ", *spaces*2)
	kelvinCelsiusOffset := 273.15

	data, err := fetchData(*token, *cityID)
	if err != nil {
		return err
	}

	description := ""
	if len(data.Weather) > 0 {
		description = data.Weather[0].Description + doubleSpacesString
	}

	currentTemperature := int(data.Main.Temp - kelvinCelsiusOffset)
	relativeHumidity := data.Main.Humidity

	sunriseTime := time.Unix(data.Sys.Sunrise, 0).Format("15:04")
	sunsetTime := time.Unix(data.Sys.Sunset, 0).Format("15:04")

	fmt.Printf("%s%s%s%d%s%s%s%d%s%s%s%s / %s\n", description, temperatureIcon, spacesString, currentTemperature, doubleSpacesString, rhIcon, spacesString, relativeHumidity, doubleSpacesString, sunIcon, spacesString, sunriseTime, sunsetTime)
	return nil
}

func fetchData(token string, cityID string) (*Response, error) {
	var response Response

	_, err := resty.R().SetQueryParams(map[string]string{
		"id":    cityID,
		"APPID": token,
	}).SetResult(&response).Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		return nil, err
	}

	return &response, nil
}
