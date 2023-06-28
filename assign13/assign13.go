package main

import "fmt"

func main() {
	var currScore int
	var totalScore [5]float64

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Game #%d: Enter the scoring total of Player %d: ", i+1, j+1)
			fmt.Scanln(&currScore)
			totalScore[i] += float64(currScore)
		}
	}
	currMax := totalScore[0]
	pIdx := 0
	for i := 1; i < 5; i++ {
		if totalScore[i] > currMax {
			currMax = totalScore[i]
			pIdx = i
		}
	}
	fmt.Printf("Player %d has the highest average score of %.2f.\n", pIdx+1, currMax/4)
}
