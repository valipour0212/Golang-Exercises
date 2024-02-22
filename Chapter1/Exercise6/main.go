package main

import "fmt"

func main() {
	var m = 3.0e-23
	var l float64 = 950
	var w float64
	var tedad float64

	fmt.Print("Enter W: ")
	fmt.Scanln(&w)

	tedad = (w * l) / m

	fmt.Println("Tedad ==> ", tedad)
}
