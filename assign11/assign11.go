package main

import "fmt"

func main() {
	var score, totalScore, cnt int
	var moreScore string

	for {
		fmt.Println("Enter a test score: ")
		_, err := fmt.Scanln(&score)
		if err != nil {
			fmt.Println(err)
			return
		}
		totalScore += score
		cnt++

		fmt.Println("Would you like to enter another score? (y/n)")
		_, err = fmt.Scanln(&moreScore)
		if err != nil {
			fmt.Println(err)
			return
		}
		if moreScore == "y" || moreScore == "Y" {
			continue
		} else if moreScore == "n" || moreScore == "N" {
			break
		} else {
			fmt.Println("Invalid input")
			return
		}
	}
	fmt.Printf("The average score is %f\n", float64(totalScore)/float64(cnt))
}
