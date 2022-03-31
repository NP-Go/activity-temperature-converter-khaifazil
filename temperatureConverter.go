package main

import "fmt"

func convertKelvin(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := (temperatureInput - 273.15) * (9/5) + 32
	resultCelsius := temperatureInput - 273.15
	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultFahrenheit := 1.8 *temperatureInput + 32
	resultKelvin := temperatureInput + 273.15
	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	//Insert your code here
	resultKelvin := (5 / 9) * (temperatureInput + 459.67)
	resultCelsius := (5 / 9) * (temperatureInput - 32)
	//Do not remove this line
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		//Insert Code here
		resultFahrenheit, resultCelsius := convertKelvin(temperatureInput)
		//DO not remove this
		fmt.Printf("Fahrenheit: %0.2f  Celsius: %0.2f", resultFahrenheit,  resultCelsius)
	} else if temperatureChoice == 2 {
		//Insert Code here
		resultFahrenheit, resultKelvin := convertCelsius(temperatureInput)
		//DO not remove this
		fmt.Printf("Fahrenheit: %0.2f  Kelvin: %0.2f", resultFahrenheit, resultKelvin)
	} else if temperatureChoice == 3 {
		//Insert Code here
		resultKelvin, resultCelsius := convertFahrenheit(temperatureInput)
		//DO not remove this
		fmt.Printf("Kelvin: %0.2f  Celsius: %0.2f", resultKelvin, resultCelsius)
	} else {
		fmt.Println("No Conversion")
	}

}
