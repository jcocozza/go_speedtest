package weather

import (
	"bytes"
	"fmt"
	"log/slog"
	"text/template"

	"github.com/jcocozza/go_speedtest/sql"
)


type WeatherSpeedTestData struct {
	SpeedTestId 		   int 					  `json:"speedtest_id"`
	CurrentWeatherResponse CurrentWeatherResponse `json:"current_weather"`
}


func LoadTemplate(weatherSpeedTestData WeatherSpeedTestData) string {
	var err error
	tmp1 := template.New("insert_template")
	tmp1, err = tmp1.Parse(sql.InsertWeatherTemplate)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}
	var buf bytes.Buffer
	err = tmp1.Execute(&buf, weatherSpeedTestData)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}
	return buf.String()
}