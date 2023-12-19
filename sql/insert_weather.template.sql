INSERT INTO weather (
    SpeedTest_id,
    LocationName, Region, Country, Lat, Lon, TzID, LocaltimeEpoch, Localtime,
    LastUpdatedEpoch, LastUpdated, TempC, TempF, IsDay,
    ConditionText, ConditionIcon, ConditionCode,
    WindMph, WindKph, WindDegree, WindDir,
    PressureMb, PressureIn,
    PrecipMm, PrecipIn,
    Humidity, Cloud, FeelslikeC, FeelslikeF, VisKm, VisMiles, Uv, GustMph, GustKph
    Co, No2, O3, So2, Pm25, Pm10, UsEpaIndex, GbDefraIndex
) VALUES (
    {{ .SpeedTestId }},
    {{ .CurrentWeatherResponse.Location.LocationName }}, {{ .CurrentWeatherResponse.Location.Region }}, {{ .CurrentWeatherResponse.Location.Country }}, {{ .CurrentWeatherResponse.Location.Lat }}, {{ .CurrentWeatherResponse.Location.Lon }}, {{ .CurrentWeatherResponse.Location.TzID }}, {{ .CurrentWeatherResponse.Location.LocaltimeEpoch }}, {{ .CurrentWeatherResponse.Location.Localtime }},
    {{ .CurrentWeatherResponse.Current.LastUpdatedEpoch }}, {{ .CurrentWeatherResponse.Current.LastUpdated }}, {{ .CurrentWeatherResponse.Current.TempC }}, {{ .CurrentWeatherResponse.Current.TempF }}, {{ .CurrentWeatherResponse.Current.IsDay }},
    {{ .CurrentWeatherResponse.Current.Condition.ConditionText }}, {{ .CurrentWeatherResponse.Current.Condition.ConditionIcon }}, {{ .CurrentWeatherResponse.Current.Condition.ConditionCode }},
    {{ .CurrentWeatherResponse.Current.WindMph}}, {{ .CurrentWeatherResponse.Current.WindKph }}, {{ .CurrentWeatherResponse.Current.WindDegree }}, {{ .CurrentWeatherResponse.Current.WindDir }},
    {{ .CurrentWeatherResponse.Current.PressureMb}}, {{ .CurrentWeatherResponse.Current.PressureIn}},
    {{ .CurrentWeatherResponse.Current.PrecipMm}}, {{ .CurrentWeatherResponse.Current.PrecipIn}},
    {{ .CurrentWeatherResponse.Current.Humidity}}, {{ .CurrentWeatherResponse.Current.Cloud}}, {{ .CurrentWeatherResponse.Current.FeelslikeC}}, {{ .CurrentWeatherResponse.Current.FeelslikeF}}, {{ .CurrentWeatherResponse.Current.VisKm}}, {{ .CurrentWeatherResponse.Current.VisMiles}}, {{ .CurrentWeatherResponse.Current.Uv}}, {{ .CurrentWeatherResponse.Current.GustMph}}, {{ .CurrentWeatherResponse.Current.GustKph}}
    {{ .CurrentWeatherResponse.Current.AirQuality.Co}}, {{ .CurrentWeatherResponse.Current.AirQuality.No2}}, {{ .CurrentWeatherResponse.Current.AirQuality.O3}}, {{ .CurrentWeatherResponse.Current.AirQuality.So2}}, {{ .CurrentWeatherResponse.Current.AirQuality.Pm25}}, {{ .CurrentWeatherResponse.Current.AirQuality.Pm10}}, {{ .CurrentWeatherResponse.Current.AirQuality.UsEpaIndex}}, {{ .CurrentWeatherResponse.Current.AirQuality.GbDefraIndex}}
)
