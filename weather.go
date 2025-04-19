package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func setCords(latitude float64, longitude float64) []string {
	if latitude > 90 || latitude < -90 {
		fmt.Println("latitude must be between 0 and 90")
		os.Exit(1)
	} else if longitude > 180 || longitude < -180 {
		fmt.Println("longitude must be between 0 and 180")
		os.Exit(1)
	}
	cords := []string{strconv.FormatFloat(latitude, 'f', 2, 64), strconv.FormatFloat(longitude, 'f', 2, 64)}
	return cords
}

//func weatherReqLatLong(api string, cords []string) *http.Response {
//	url := "https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s"
//	url = fmt.Sprintf(url, cords[0], cords[1], api)
//	weather, _ := http.Get(url)
//
//	return weather
//}

func weatherReqCity(api string, city string) WeatherResponse {
	url := "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial"
	url = fmt.Sprintf(url, city, api)
	weather, _ := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(weather.Body)

	var result WeatherResponse
	err := json.NewDecoder(weather.Body).Decode(&result)
	if err != nil {
		return WeatherResponse{}
	}
	return result
}

func main() {
	var key = os.Getenv("WEATHER_KEY")
	fmt.Println("WEATHER_KEY:", key)
	test := setCords(10, 20)
	fmt.Println(test)
	//fmt.Println(weatherReqLatLong(key, test))
	fmt.Println(weatherReqCity(key, "Boise"))
}
