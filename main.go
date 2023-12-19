package main

import (
	//"github.com/jcocozza/go_speedtest/database"
	"fmt"

	//"github.com/jcocozza/go_speedtest/speedtest"
	"github.com/jcocozza/go_speedtest/weather"
)

const dbFileName = "go-speedtest.db"

func main() {
	/*
	db, err := database.Connect(dbFileName)
	if err != nil {
		panic(err)
	}
	*/

	//result := speedtest.SpeedTestJSON()
	//fmt.Println(result)
	result := weather.GetWeather()
	//fmt.Println(result)

	wstd := weather.WeatherSpeedTestData{
		SpeedTestId: 1,
		CurrentWeatherResponse: result,
	}

	fmt.Println(weather.LoadTemplate(wstd))
}