package datatype

func CelsiusToKelvin(in float64) float64 {
	return in + 273.15
}
func CelsiusToFahrenheit(in float64) float64 {
	return (in - 32) / 1.8
}

func KelvinToCelsius(in float64) float64 {
	return in - 273.15
}
func KelvinToFahrenheit(in float64) float64 {
	return (in-273.15)*1.8 + 32
}

func FahrenheitToKelvin(in float64) float64 {
	return (in-32)/1.8 + 273.15
}
func FahrenheitToCelsius(in float64) float64 {
	return (in - 32) / 1.8
}
