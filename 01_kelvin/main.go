package main

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
)

func main() {
	var todayF Fahrenheit = 48
	var todayC Celsius
	var todayK Kelvin

	fmt.Println(CtoF(0))   // should be 32
	fmt.Println(CtoF(100)) // should be 212
	todayC = FtoC(todayF)
	fmt.Println(todayC)

	fmt.Println(FtoK(32))
	fmt.Println(FtoK(212))
	todayK = FtoK(todayF)
	fmt.Println(todayK)

	fmt.Println(CtoK(0))
	fmt.Println(CtoK(100))
	todayK = CtoK(todayC)
	fmt.Println(todayK)
}

// CelsiusToFahrenheit converts celsius to fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(float64((c * 9 / 5) + 32))
}

// FahrenheitToCelsius converts fahrenheit to celsius.
func FtoC(f Fahrenheit) Celsius {
	return Celsius(float64((f - 32) * 5 / 9))
}

// FahrenheitToKelvin
func FtoK(f Fahrenheit) Kelvin {
	return Kelvin(float64((f + 459.67) * 5 / 9))
}

// CelsiustoKelvin
func CtoK(c Celsius) Kelvin {
	return Kelvin(float64(c + 273.15))
}
