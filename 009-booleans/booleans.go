package main

import (
	"fmt"
	"strings"
)

func main() {
	columns := []bool{true, false}
	rows := []bool{true, false}

	fmt.Println("And Table")
	header := fmt.Sprintf("|%3s    | %3s    | p && q |", "p", "q")
	header = header + "\n" + strings.Repeat("-", len(header))
	fmt.Println(header)

	for _, p := range columns {
		for _, q := range rows {
			fmt.Printf("|%6t | %6t | %6t |\n", p, q, p && q)
		}
	}

	fmt.Println("\nOr Table")
	header = fmt.Sprintf("|%3s    | %3s    | p || q |", "p", "q")
	header = header + "\n" + strings.Repeat("-", len(header))
	fmt.Println(header)

	for _, p := range columns {
		for _, q := range rows {
			fmt.Printf("|%6t | %6t | %6t |\n", p, q, p || q)
		}
	}
}
