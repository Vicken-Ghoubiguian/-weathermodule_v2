package weathermodule_v2

type Sys struct {
	Type int64
	Id int64
	Country string
	Sunrise int64
	Sunset int64
}

type Main struct {
	Temp float64
	Feels_like float64
	Temp_min float64
	Temp_max float64
	Pressure float64
	Humidity int64
}

type WindStruct struct {
	Speed float64
	Deg int64
	Gust float64
}

type CloudsStruct struct {
	All int64
}

type CoordinatesStruct struct {
	Lon float64
	Lat float64
}

type WeatherStruct struct {
	Id int64
	Main string
	Description string
	Icon string
}

type OWMStruct struct {
	Coord CoordinatesStruct `json:"coord"`
	Weather [1]WeatherStruct `json:"weather"`
	Base string `json:"base"`
	Main Main `json:"main"`
	Visibility int64 `json:"visibility"`
	Wind WindStruct `json:"wind"`
	Clouds CloudsStruct `json:"clouds"`
	Dt int64 `json:"dt"`
	Sys Sys `json:"sys"`
	Timezone int64 `json:"timezone"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	Cod int64 `json:"cod"`
	Message string `json:"message"`
}