package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var pair, vPair, howMany int
	maxRollings := 100000

	fmt.Println("Reading pair of dice value...")
	fmt.Scanf("%d", &vPair)

	for i := 0; i < maxRollings; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2

		if pair == vPair {
			howMany++
		}
	}

	fmt.Printf("Probability of rolling a %d is %4g\n", vPair, float32(howMany)/float32(maxRollings))
}
