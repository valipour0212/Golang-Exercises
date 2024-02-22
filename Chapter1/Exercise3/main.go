package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.14

	var radius float64

	fmt.Print("Radius of Sphere: ")
	fmt.Scanln(&radius)

	surArea := 4 * PI * math.Pow(radius, 2)
	volume := (4 / 3) * (PI * math.Pow(radius, 3))

	fmt.Println("Surface Area is: ", surArea)
	fmt.Println("Volume is: ", volume)
}
