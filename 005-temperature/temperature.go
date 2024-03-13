package main

import "fmt"

func main() {
	var from = 0
	var to = 250
	var by = 10
	var fahrenheit = float32(from)
	var centigrade float32

	fmt.Println("|   Fahrenheit\t| Centigrade\t|")

	for from <= to {
		centigrade = 5.0/9.0 * (fahrenheit - 32)

		fmt.Printf("|\t%g ºF\t|  %.2f ºC\t|\n", fahrenheit, centigrade)

		fahrenheit += float32(by)
		from += by
	}
}