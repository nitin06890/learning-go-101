package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nFlips, heads, tails int

	fmt.Println("How many times would you like to flip the coin?")
	fmt.Scanln(&nFlips)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nFlips; i++ {
		result := rand.Intn(2)
		if result == 0 {
			heads++
		} else {
			tails++
		}
	}

	fmt.Printf("After flipping the coin %d times, there were %d heads and %d tails.\n", nFlips, heads, tails)
}
