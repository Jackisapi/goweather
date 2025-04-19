package main

import (
	"flag"
	"fmt"
	weather "goweather/libs"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	var city string
	flag.StringVar(&city, "city", "Boise", "Enter a City Name")

	var countryCode string
	flag.StringVar(&countryCode, "country", "US", "Enter a Country")

	flag.Parse()

	var key = os.Getenv("WEATHER_KEY")
	if key == "" {
		panic("Missing WEATHER_KEY")
		os.Exit(1)
	}
	fmt.Println(weather.WeatherReqCity(key, city, countryCode))

}
