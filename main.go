package main

import (
	"encoding/json"
	"fmt"
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

func main() {

	godotenv.Load()
	var lat, lon float64
	fmt.Printf("Enter the lat and lon :")
	fmt.Scanf("%v %v", &lat, &lon)

	client := &http.Client{}

	ApiKey := os.Getenv("OpenWeatherMapApiKey")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v", lat, lon, ApiKey), nil)
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
	fmt.Println(forecast)

}
