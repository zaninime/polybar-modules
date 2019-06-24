package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/xerrors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/resty.v1"
)

var (
	token  = kingpin.Flag("token", "API token for aqicn.org.").Short('t').Required().String()
	city   = kingpin.Flag("city", "City to fetch data from. Defaults to current location (IP-based).").String()
	icon   = kingpin.Flag("icon", "Icon used as a prefix").Default("\uf299").String()
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
	feedName, err := getFeedName(*city)
	if err != nil {
		return xerrors.Errorf("error while retrieving feed name: %w", err)
	}

	data, err := fetchData(feedName, *token)
	if err != nil {
		return xerrors.Errorf("error while fetching aqi API: %w", err)
	}

	aqi := data.Data.Aqi
	spacesString := strings.Repeat(" ", *spaces)
	var color string

	switch {
	case aqi < 50:
		color = "#009966"
	case aqi < 50:
		color = "#ffde33"
	case aqi < 50:
		color = "#ff9933"
	case aqi < 50:
		color = "#660099"
	default:
		color = "#7e0023"
	}

	fmt.Printf("%%{F%s}%s%%{F-}%s%d\n", color, *icon, spacesString, aqi)
	return nil
}

func getFeedName(city string) (string, error) {
	if city != "" {
		return city, nil
	}

	var response struct {
		Location struct {
			Lat float32 `json:"lat"`
			Lng float32 `json:"lng"`
		} `json:"location"`
	}

	_, err := resty.R().SetResult(&response).Get("https://location.services.mozilla.com/v1/geolocate?key=geoclue")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("geo:%f;%f", response.Location.Lat, response.Location.Lng), nil
}

func fetchData(feedName string, token string) (*Response, error) {
	var responseData Response
	_, err := resty.R().SetPathParams(map[string]string{
		"feedName": feedName,
	}).SetQueryParam("token", token).SetResult(&responseData).Get("https://api.waqi.info/feed/{feedName}/")

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
