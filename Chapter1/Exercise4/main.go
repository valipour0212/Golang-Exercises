package main

import (
	"fmt"
	"math"
)

func main() {
	var nSides, sLength float64

	fmt.Print("Enter Number of Sides: ")
	fmt.Scanln(&nSides)

	fmt.Print("Enter The Length of a Side: ")
	fmt.Scanln(&sLength)

	pArea := (nSides * math.Pow(sLength, 2)) / (4 * math.Tan(math.Pi/nSides))

	fmt.Println("The Area of the Polygon is: ", pArea)
}
