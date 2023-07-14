package main

// Funciones de conversion de temperaturas

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
