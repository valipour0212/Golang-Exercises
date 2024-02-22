package main

import (
	"fmt"
	"math"
)

func main() {
	var v, t float64

	fmt.Print("Enter wind speed in Km/hour: ")
	fmt.Scanln(&v)

	fmt.Print("Enter air temperature in degrees Celsius: ")
	fmt.Scanln(&t)

	wci := 13.12 + 0.6215*t - 11.37*math.Pow(v, 0.16) + 0.3965*t*math.Pow(v, 0.16)

	fmt.Println("The wind chill index is: ", math.Round(wci))
}
