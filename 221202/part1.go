package main

import (
	"fmt"
	"os"
	"strings"
)

func getScore(enemy string, me string) int {
	fmt.Println("Selections: ", enemy, me)

	shapeScore := 0
	if me == "X" {
		shapeScore = 1
	} else if me == "Y" {
		shapeScore = 2
	} else if me == "Z" {
		shapeScore = 3
	}

	roundScore := 0
	if (me == "X" && enemy == "A") || (me == "Y" && enemy == "B") || (me == "Z" && enemy == "C") {
		roundScore = 3
	} else if (me == "X" && enemy == "C") || (me == "Y" && enemy == "A") || (me == "Z" && enemy == "B") {
		roundScore = 6
	}
	score := shapeScore + roundScore
	fmt.Println(score, shapeScore, roundScore)
	return score
}

func Part1() {
	//myFile, _ := os.ReadFile("221202/test_input.txt")
	myFile, _ := os.ReadFile("221202/input.txt")

	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")

	score := 0
	for _, s := range lines {
		battle := strings.Split(s, " ")

		score += getScore(battle[0], battle[1])
		//if s == "" {
		//	if cur > highest {
		//		highest = cur
		//	}
		//	cur = 0
		//} else {
		//	sval, _ := strconv.Atoi(s)
		//	cur += sval
		//}
	}
	//if cur > highest {
	//	highest = cur
	//}
	fmt.Println(score)
}
