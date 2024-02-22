package main

import "fmt"

func main() {
	const PI = 3.14
	var height, radius float32

	fmt.Print("Height of Cylinder: ")
	fmt.Scanln(&height)

	fmt.Print("Radius of Cylinder: ")
	fmt.Scanln(&radius)

	volume := int(PI * radius * radius * height)
	surArea := int((2 * PI * radius * height) + (2 * PI * radius * radius))

	fmt.Println("Volume is: ", volume)
	fmt.Println("Surface is: ", surArea)
}
