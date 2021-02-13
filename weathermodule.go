package weathermodule_v2

// Import all necessary packages
import (
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Defining the type 'Weather' which recover and manage current weather in a defined city
type WeatherModule struct {

	//
	coords Coordinates

	//
	weather Weather

	//
	temperature Temperature
	feelingLikeTemperature Temperature
	minTemperature Temperature
	maxTemperature Temperature

	//
	pressure Pressure

	//
	humidity Humidity

	//
	wind Wind

	//
	sunrise SunTime
	sunset SunTime

	//
	geographicLocation GeographicLocation

	//
	ultraViolet UV
}

// Defining the Weather initializer
func (w *WeatherModule) InitializeWeatherModule(city string, countrysISOAlpha2Code string, apiKey string) {

	//
	varOWMStruct

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city + "," + countrysISOAlpha2Code, apiKey)

	//
	weatherResp, err := http.Get(weatherRequest)

	//
	OtherErrorHandlerFunction(err, Red(), Reset())

	//
	weatherJsonString, err := ioutil.ReadAll(weatherResp.Body)

	//
	OtherErrorHandlerFunction(err, Red(), Reset())

	//Single instruction to convert weather_json_string []byte variable to string
	err = json.Unmarshal(weatherJsonString, &owm)

	//
	OtherErrorHandlerFunction(err, Red(), Reset())

	//
	if owm.Cod != 200 {

		// Calling the 'owmErrorHandlerFunction' from the module to treat the current error...
		OwmErrorHandlerFunction(strconv.Itoa(int(owm.Cod)), Red(), owm.Message, Reset())

	} else {

		//
		var UVowm UVStruct

		//
		uvRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, fmt.Sprintf("%g", owm.Coord.Lat), fmt.Sprintf("%g", owm.Coord.Lon))

		//
		uvResp, err := http.Get(uvRequest)

		//
		OtherErrorHandlerFunction(err, Red(), Reset())

		//
		uvJsonString, err := ioutil.ReadAll(uvResp.Body)

		//
		OtherErrorHandlerFunction(err, Red(), Reset())

		//
		err = json.Unmarshal(uvJsonString, &UVowm)

		//
		OtherErrorHandlerFunction(err, Red(), Reset())

		//
		w.coords.InitializeCoordinates(owm.Coord.Lon, owm.Coord.Lat)

		//
		w.weather.InitializeWeather(owm.Weather[0].Id, owm.Weather[0].Main, owm.Weather[0].Description, owm.Weather[0].Icon)

		//
		w.temperature.InitializeTemperature(owm.Main.Temp)
		w.feelingLikeTemperature.InitializeTemperature(owm.Main.Feels_like)
		w.minTemperature.InitializeTemperature(owm.Main.Temp_min)
		w.maxTemperature.InitializeTemperature(owm.Main.Temp_max)

		//
		w.wind.InitializeWind(owm.Wind.Speed, owm.Wind.Deg, owm.Wind.Gust)

		//
		w.pressure.InitializePressure(owm.Main.Pressure)

		//
		w.humidity.InitializeHumidity(owm.Main.Humidity)

		//
		w.geographicLocation.InitializeGeographicLocation(owm.Sys.Country, owm.Name)

		//
		w.sunrise.InitializeSunTime(owm.Sys.Sunrise)
		w.sunset.InitializeSunTime(owm.Sys.Sunset)

		//
		w.ultraViolet.InitializeUV(int64(UVowm.Value))

		// Displaying success message...
		fmt.Println(Green() + "Weather implemented successfully !" + Reset() + "\n")
	}
}

// Defining the minimal Weather initializer
func (w *WeatherModule) InitializeMinimallyWeatherModule(city string, apiKey string) {

	//
	varOWMStruct

	// Defining the HTTP request's URL for weather and uv
	weatherRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	//
	weatherResp, err := http.Get(weatherRequest)

	//
	OtherErrorHandlerFunction(err, Red(), Reset())

	//
	weatherJsonString, err := ioutil.ReadAll(weatherResp.Body)

	//
	OtherErrorHandlerFunction(err, Red(), Reset())

	//Single instruction to convert weather_json_string []byte variable to string
	err = json.Unmarshal(weatherJsonString, &owm)

	//
	OtherErrorHandlerFunction(err, Red(), Reset())

	//
	if owm.Cod != 200 {

		// Calling the 'owmErrorHandlerFunction' from the ' module to treat the current error...
		OwmErrorHandlerFunction(strconv.Itoa(int(owm.Cod)), Red(), owm.Message, Reset())

	} else {

		//
		var UVowm UVStruct

		//
		uvRequest := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", apiKey, fmt.Sprintf("%g", owm.Coord.Lat), fmt.Sprintf("%g", owm.Coord.Lon))

		//
		uvResp, err := http.Get(uvRequest)

		//
		OtherErrorHandlerFunction(err, Red(), Reset())

		//
		uvJsonString, err := ioutil.ReadAll(uvResp.Body)

		//
		OtherErrorHandlerFunction(err, Red(), Reset())

		//
		err = json.Unmarshal(uvJsonString, &UVowm)

		//
		OtherErrorHandlerFunction(err, Red(), Reset())

		//
		w.coords.InitializeCoordinates(owm.Coord.Lon, owm.Coord.Lat)

		//
		w.weather.InitializeWeather(owm.Weather[0].Id, owm.Weather[0].Main, owm.Weather[0].Description, owm.Weather[0].Icon)

		//
		w.temperature.InitializeTemperature(owm.Main.Temp)
		w.feelingLikeTemperature.InitializeTemperature(owm.Main.Feels_like)
		w.minTemperature.InitializeTemperature(owm.Main.Temp_min)
		w.maxTemperature.InitializeTemperature(owm.Main.Temp_max)

		//
		w.wind.InitializeWind(owm.Wind.Speed, owm.Wind.Deg, owm.Wind.Gust)

		//
		w.pressure.InitializePressure(owm.Main.Pressure)

		//
		w.humidity.InitializeHumidity(owm.Main.Humidity)

		//
		w.geographicLocation.InitializeGeographicLocation(owm.Sys.Country, owm.Name)

		//
		w.sunrise.InitializeSunTime(owm.Sys.Sunrise)
		w.sunset.InitializeSunTime(owm.Sys.Sunset)

		//
		w.ultraViolet.InitializeUV(int64(UVowm.Value))

		// Displaying success message...
		fmt.Println(Green() + "Weather implemented successfully !" + Reset() + "\n")
	}
}

//
func (w *WeatherModule) GetCoords() Coordinates {

	return &w.coords
}

//
func (w *WeatherModule) GetWeather() Weather {

	return &w.weather
}

//
func (w *WeatherModule) GetTemperature() Temperature {

	return &w.temperature
}

//
func (w *WeatherModule) GetFeelingLikeTemperature() Temperature {

	return &w.feelingLikeTemperature
}

//
func (w *WeatherModule) GetMinTemperature() Temperature {

	return &w.minTemperature
}

//
func (w *WeatherModule) GetMaxTemperature() Temperature {

	return &w.maxTemperature
}

//
func (w *WeatherModule) GetPressure() Pressure {

	return &w.pressure
}

//
func (w *WeatherModule) GetHumidity() Humidity {

	return &w.humidity
}

//
func (w *WeatherModule) GetWind() Wind {

	return &w.wind
}

//
func (w *WeatherModule) GetSunrise() SunTime {

	return &w.sunrise
}

//
func (w *WeatherModule) GetSunset() SunTime {

	return &w.sunset
}

//
func (w *WeatherModule) GetGeographicLocation() GeographicLocation {

	return &w.geographicLocation
}

//
func (w *WeatherModule) GetUltraViolet() UV {

	return &w.ultraViolet
}
