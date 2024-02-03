package main

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}
type weatherdata struct {
	name string
	Main struct {
		Kelvin float64
	}
}

func main() {

}
