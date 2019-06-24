package main

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Aqi          int64         `json:"aqi"`
	Idx          int64         `json:"idx"`
	Attributions []Attribution `json:"attributions"`
	City         City          `json:"city"`
	Dominentpol  string        `json:"dominentpol"`
	Iaqi         Iaqi          `json:"iaqi"`
	Time         Time          `json:"time"`
	Debug        Debug         `json:"debug"`
}

type Attribution struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type City struct {
	Geo  []float64 `json:"geo"`
	Name string    `json:"name"`
	URL  string    `json:"url"`
}

type Debug struct {
	Sync string `json:"sync"`
}

type Iaqi struct {
	Co   Co `json:"co"`
	H    Co `json:"h"`
	No2  Co `json:"no2"`
	O3   Co `json:"o3"`
	P    Co `json:"p"`
	Pm10 Co `json:"pm10"`
	T    Co `json:"t"`
}

type Co struct {
	V float64 `json:"v"`
}

type Time struct {
	S  string `json:"s"`
	Tz string `json:"tz"`
	V  int64  `json:"v"`
}
