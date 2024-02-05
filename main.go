package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//	type apiConfigData struct {
//		OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
//	}
type weatherdata struct {
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}
type locadata struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func main() {

	godotenv.Load()

	var city string
	fmt.Printf("Enter the city :")
	fmt.Scanf("%v ", &city)

	client := &http.Client{}

	ApiKey := os.Getenv("OpenWeatherMapApiKey")
	locreq, err := http.NewRequest("GET", fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=2&appid=%v", city, ApiKey), nil)
	if err != nil {
		log.Println("Error in getting location", err)
	}
	locres, err := client.Do(locreq)
	if err != nil {
		fmt.Println("Error in response:", err)
		return
	}
	defer locres.Body.Close()
	location := []locadata{}
	err = json.NewDecoder(locres.Body).Decode(&location)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v", location[0].Lat, location[0].Lon, ApiKey), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	defer res.Body.Close()

	forecast := weatherdata{}
	err = json.NewDecoder(res.Body).Decode(&forecast)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	fmt.Println(forecast.Main.Kelvin)

}
