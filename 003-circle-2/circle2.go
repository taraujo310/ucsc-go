package main

import (
	"fmt"
	"math"
)

func main() {
	var area, radius float32

	fmt.Print("Input radius: ")
	fmt.Scanf("%f", &radius)

	area = math.Pi * radius * radius

	fmt.Printf("Area = %f\n", area)
}