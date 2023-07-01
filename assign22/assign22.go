package main

import "fmt"

type Player struct {
	totalPoints int
	totalGames  int
}

func main() {
	players := [5]Player{}

	for i := 0; i < 5; i++ {
		fmt.Println("Enter player", i+1, "total points:")
		fmt.Scanln(&players[i].totalPoints)
		fmt.Println("Enter player", i+1, "total games:")
		fmt.Scanln(&players[i].totalGames)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Player %d scoring average: %.2f\n", i+1, float64(players[i].totalPoints)/float64(players[i].totalGames))
	}
}
