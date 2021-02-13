# weatherModule_v2

This is the second (and final) version of a Go package using OpenWeatherMap to get and manipulate current weather datas from a wished city in a wished country.

## Contents

1. [What is this project ?](#what_is_this_project)

2. [How to use it ?](#how_to_use_it)

3. [A little example](#a_little_example)

4. [A few usefull links](#a_few_usefull_links)

5. [Conclusion](#conclusion)

<a name="what_is_this_project"></a>
## What is this project ?

This project is about the development of a Go package using OpenWeatherMap to get and manipulate current weather datas from a wished city in a wished country.
This repository is composed of one Go file named *weathermodule.go* which contain the *WeatherModule* main type to get, to encapsulate and to manage weather, one Dockerfile to build your own Docker image to run examples, one *LICENCE* file to define the license (here the GNU version 3.0 one), the [*src*](https://github.com/Vicken-Ghoubiguian/weathermodule/tree/master/src) directory which contains all defined classes used in the WeatherModule type definition and the tthe [*samples*](https://github.com/Vicken-Ghoubiguian/weathermodule/tree/master/samples) directory which contain all defined samples to test locally the *WeatherModule* type, all of the other types used to define weather and their functions.

<a name="how_to_use_it"></a>
## How to use it ?

```bash
$ go get -u github.com/Vicken-Ghoubiguian/weathermodule_v2
```
This will retrieve the library.

<a name="a_little_example"></a>
## A little example

```go
package main

// import all necessary packages
import (
	"github.com/Vicken-Ghoubiguian/weathermodule_v2"
	"fmt"
)

// main function to test all of the package
func main() {

	//
	var weatherObj weathermodule_v2.WeatherModule

	weatherObj.InitializeWeatherModule("<wished_city_name>", "<wished_country_code>", "<Your_OpenWeather_Map_API_key>")

	fmt.Printf("" + weatherObj.GetGeographicLocation().GetCityName() + " (" + weatherObj.GetGeographicLocation().GetCountryCode() + ")\n")

	fmt.Printf("Weather (" + weatherObj.GetWeather().GetMain() + ", " + weatherObj.GetWeather().GetDescription() + ", " + weatherObj.GetWeather().GetIconUrl() + ")\n")

	fmt.Printf("(" + fmt.Sprintf("%f", weatherObj.GetCoords().GetLongitude()) + ", " + fmt.Sprintf("%f", weatherObj.GetCoords().GetLatitude()) + ")\n")

	fmt.Printf("\n")

	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsCelsius()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsFahrenheit()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	weatherObj.GetTemperature().SetTemperatureAsKelvin()
	fmt.Printf("Temperature (in " + weatherObj.GetTemperature().GetCurrentTemperatureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetTemperature().GetTemperatureValue()) + weatherObj.GetTemperature().GetTemperatureScaleSymbol() + "\n")

	fmt.Printf("\n")

	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + " " + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	weatherObj.GetPressure().SetPressureAsPascal()
	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	weatherObj.GetPressure().SetPressureAsBar()
	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	weatherObj.GetPressure().SetPressureAsAtmosphere()
	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	weatherObj.GetPressure().SetPressureAsTorr()
	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	weatherObj.GetPressure().SetPressureAsPoundsPerSquareInch()
	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	weatherObj.GetPressure().SetPressureAsHectoPascal()
	fmt.Printf("Pressure (in " + weatherObj.GetPressure().GetPressureScale().String() + "): " + fmt.Sprintf("%f", weatherObj.GetPressure().GetPressureValue()) + weatherObj.GetPressure().GetPressureSymbolUnit() + "\n")

	fmt.Printf("\n")

	fmt.Printf("UV index: " + fmt.Sprintf("%d", weatherObj.GetUltraViolet().GetIndex()) + ", UV risk: " + weatherObj.GetUltraViolet().GetRisk().String() + "\n")

	fmt.Printf("Humidity: " + fmt.Sprintf("%d", weatherObj.GetHumidity().GetHumidityValue()) + " " + weatherObj.GetHumidity().GetHumidityUnitScale() + "\n")

	fmt.Printf("Wind speed: " + fmt.Sprintf("%f", weatherObj.GetWind().GetSpeed()) + "\n")
	fmt.Printf("Wind Deg: " + fmt.Sprintf("%d", weatherObj.GetWind().GetDeg()) + "\n")
	fmt.Printf("Wind Gust: " + fmt.Sprintf("%f", weatherObj.GetWind().GetGust()) + "\n")

	weatherObj.GetSunrise().SetCurrentFormatAsTimestamp()
	weatherObj.GetSunset().SetCurrentFormatAsTimestamp()

	fmt.Printf("Sunrise (as " + weatherObj.GetSunrise().GetCurrentFormat().String() + "): " + weatherObj.GetSunrise().GetSunTimeInCurrentFormat() + "\n")
	fmt.Printf("Sunset (as " + weatherObj.GetSunset().GetCurrentFormat().String() + "): " + weatherObj.GetSunset().GetSunTimeInCurrentFormat() + "\n")
}
```

<a name="a_few_usefull_links"></a>
## A few usefull links

<a name="conclusion"></a>
## Conclusion

It has been an exciting project, easily usable by anyone in any other project written in Go language regardless of its size (large or small) and which has broadened my knowledge in Go language development as potentially every other beginner or confirmed developer.
