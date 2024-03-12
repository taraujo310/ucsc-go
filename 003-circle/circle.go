package main

import "fmt"

func main() {
	var area, radius float32
	const pi = 3.14159

	fmt.Print("Input radius: ")
	fmt.Scanf("%f", &radius)

	area = pi * radius * radius

	fmt.Printf("Area = %f\n", area)
}