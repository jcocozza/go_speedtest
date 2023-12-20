package main

import (
	"fmt"
	"log/slog"

	"github.com/jcocozza/go_speedtest/database"

	"github.com/jcocozza/go_speedtest/speedtest"
	"github.com/jcocozza/go_speedtest/weather"
)

func main() {
	db := database.ConnectToDatabase()

	speedTestResult := speedtest.SpeedTestJSON()
	weatherResult := weather.GetWeather()

	insertSpeedTestSQL := speedtest.LoadTemplate(speedTestResult)

	result, err := db.Exec(insertSpeedTestSQL)
	if err != nil {
		slog.Error("unable to insert speedtest sql: " + err.Error())
		panic(err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("Unable to get last insert id: " + fmt.Sprint(err))
		panic(err)
	}
	lastIDint := int(lastID)

	wstd := weather.WeatherSpeedTestData{
		SpeedTestId: lastIDint,
		CurrentWeatherResponse: weatherResult,
	}

	insertWeatherSQL := weather.LoadTemplate(wstd)

	_, err = db.Exec(insertWeatherSQL)
	if err != nil {
		slog.Error("Unable to get insert weather sql: " + fmt.Sprint(err))
	}
}