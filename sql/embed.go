package sql

import _ "embed"

//go:embed init.sql
var SqlInit string

//go:embed insert_weather.template.sql
var InsertWeatherTemplate string

//go:embed insert_speedtest.template.sql
var InsertSpeedTestTemplate string