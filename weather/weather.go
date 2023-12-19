package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	BASEURL = "http://api.weatherapi.com/v1"
)

// get the lat/long pair for the machine's ip address
func getLocation() (float64, float64) {
	cmd := exec.Command("curl", "ipinfo.io", "-s")
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Error: " + fmt.Sprint(err))
		panic(err)
	}
	var ipInfo IPInfoResponse
	err = json.Unmarshal(output, &ipInfo)
    if err != nil {
        slog.Error("Error unmarshaling JSON: " + fmt.Sprint(err))
        panic(err)
    }

	coordsStrLst := strings.Split(ipInfo.Loc, ",")
	latitude, err := strconv.ParseFloat(coordsStrLst[0], 64)
	longitude, err := strconv.ParseFloat(coordsStrLst[1], 64)

	if err != nil {
		slog.Error("Error parsing coordinates:", fmt.Sprint(err))
	}

	return latitude, longitude
}

// generate the api string for current weather data
func genCurrentUrl(key string, latitude, longitude float64) string {
	currentUrl := BASEURL + "/" + "current.json?key=" + key + "&q=" + fmt.Sprint(latitude) + "," + fmt.Sprint(longitude) + "&aqi=yes"
	slog.Info("Weather URL: " + currentUrl)
	return currentUrl
}

// function to get the current weather conditions
func GetWeather() CurrentWeatherResponse {
	var latitude float64
	var longitude float64
	var err error

	// handle lat/long coords
	if len(os.Args) > 1 {
		latitudeStr, longitudeStr := os.Args[0], os.Args[1]

		latitude, err = strconv.ParseFloat(latitudeStr, 64)
		if err != nil {
			slog.Error("Error parsing latitude: " + fmt.Sprint(err))
		}

		longitude, err = strconv.ParseFloat(longitudeStr, 64)
		if err != nil {
			slog.Error("Error parsing longitude: " + fmt.Sprint(err))
		}
	} else {
		latitude, longitude = getLocation()
	}

	url := genCurrentUrl(APIKEY, latitude, longitude)

	response, err := http.Get(url)
	var data []byte
	if err != nil {
		slog.Error("Failed to get weather: " + fmt.Sprint(err))
		panic(err)
	} else {
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			slog.Error("Failed to read response body: " + fmt.Sprint(err))
			panic(err)
		}
		data = responseBody
		slog.Info("weather response:\n" + string(data))
	}

	var currentWeatherResponse CurrentWeatherResponse
	err = json.Unmarshal(data, &currentWeatherResponse)
    if err != nil {
        slog.Error("Error unmarshaling JSON: " + fmt.Sprint(err))
        panic(err)
    }
	return currentWeatherResponse

}