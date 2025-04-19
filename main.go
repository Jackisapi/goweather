package main

import (
	"fmt"
	"os"
	"screwinoff/libs"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var key = os.Getenv("WEATHER_KEY")
	//fmt.Println(weatherReqLatLong(key, test))
	boiseWeather := weather.WeatherReqCity(key, "Boise", "us")
	fmt.Println(boiseWeather)
	fmt.Println("\n")
	tokyoWeather := weather.WeatherReqCity(key, "Tokyo", "jp")
	fmt.Println(tokyoWeather)

}
