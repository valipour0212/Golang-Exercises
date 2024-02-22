package main

import "fmt"

func main() {
	var base, height float32

	fmt.Print("Length of Base: ")
	fmt.Scanln(&base)

	fmt.Print("Measurement of Height: ")
	fmt.Scanln(&height)

	area := base * height

	print("Area is: ", int(area))
}
