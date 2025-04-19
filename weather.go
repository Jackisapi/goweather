package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// sets how the data is to be contained
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

// Sets the cords into a list
func setCords(latitude float64, longitude float64) []string {
	// checks to see if latitude is beyond range
	if latitude > 90 || latitude < -90 {
		fmt.Println("latitude must be between 0 and 90")
		//exits if so
		os.Exit(1)
		//checks to see if longitude is beyond range
	} else if longitude > 180 || longitude < -180 {
		fmt.Println("longitude must be between 0 and 180")
		//exits if so
		os.Exit(1)
	}
	//converts the cords into a string and puts them into a list
	cords := []string{strconv.FormatFloat(latitude, 'f', 2, 64), strconv.FormatFloat(longitude, 'f', 2, 64)}
	// Finally returns the value
	return cords
}

//This does not work lol
//func weatherReqLatLong(api string, cords []string) *http.Response {
//	url := "https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s"
//	url = fmt.Sprintf(url, cords[0], cords[1], api)
//	weather, _ := http.Get(url)
//
//	return weather
//}

func weatherReqCity(api string, city string) WeatherResponse {
	url := "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial"
	//formats the request url with the city, and api key
	url = fmt.Sprintf(url, city, api)
	//makes the request
	weather, _ := http.Get(url)
	// waits until request pings back to continue running errors if it doesnt
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(weather.Body)

	//creates json template using the scruct built above
	var result WeatherResponse
	//encodes the response into the above format and returns
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
