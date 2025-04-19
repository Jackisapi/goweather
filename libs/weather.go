package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WeatherResponse sets how the data is to be contained
type WeatherResponse struct {
	Name     string `json:"name"`
	Timezone int    `json:"timezone"`
	DT       int64  `json:"dt"`
	Main     struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
}

// Assures that when the struct is printed it doest look trashy
func (w WeatherResponse) String() string {
	t := time.Unix(w.DT, 0).UTC().Add(time.Second * time.Duration(w.Timezone))
	return fmt.Sprintf("City: %s\nTime: %s\nTemp: %.1f°C\nFeels Like: %.1f°C\nHumidity: %d%% \nWind: (Speed: %f, Angle: %d) \n Desc: %s",
		w.Name,
		t.Format("2006-01-02 15:04:05"),
		w.Main.Temp,
		w.Main.FeelsLike,
		w.Main.Humidity,
		w.Wind.Speed,
		w.Wind.Deg,
		w.Weather[0].Description,
	)
}

//// Sets the cords into a list
//func setCords(latitude float64, longitude float64) []string {
//	// checks to see if latitude is beyond range
//	if latitude > 90 || latitude < -90 {
//		fmt.Println("latitude must be between 0 and 90")
//		//exits if so
//		os.Exit(1)
//		//checks to see if longitude is beyond range
//	} else if longitude > 180 || longitude < -180 {
//		fmt.Println("longitude must be between 0 and 180")
//		//exits if so
//		os.Exit(1)
//	}
//	//converts the cords into a string and puts them into a list
//	cords := []string{strconv.FormatFloat(latitude, 'f', 2, 64), strconv.FormatFloat(longitude, 'f', 2, 64)}
//	// Finally returns the value
//	return cords
//}

//This does not work lol
//func weatherReqLatLong(api string, cords []string) *http.Response {
//	url := "https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s"
//	url = fmt.Sprintf(url, cords[0], cords[1], api)
//	weather, _ := http.Get(url)
//
//	return weather
//}

func WeatherReqCity(api string, city string, country_code string) WeatherResponse {
	url := "https://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s&units=imperial"
	//formats the request url with the city, and api key
	url = fmt.Sprintf(url, city, country_code, api)
	//makes the request
	weather, _ := http.Get(url)
	// waits until request pings back to continue running errors if it doesn't
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(weather.Body)

	//creates json template using the struct built above
	var result WeatherResponse
	//encodes the response into the above format and returns
	data := json.NewDecoder(weather.Body).Decode(&result)
	if data != nil {
		return WeatherResponse{}
	}
	return result
}
