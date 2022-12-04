package main

import (
	"fmt"
	"os"
	"strings"
)

func getMove(enemy string, outcome string) int {

	xLookup := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	yLookup := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	zLookup := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	me := ""
	if outcome == "X" {
		me = xLookup[enemy]
	} else if outcome == "Y" {
		me = yLookup[enemy]
	} else if outcome == "Z" {
		me = zLookup[enemy]
	}

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

func Part2() {
	//myFile, _ := os.ReadFile("221202/test_input.txt")
	myFile, _ := os.ReadFile("221202/input.txt")

	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")

	score := 0
	for _, s := range lines {
		battle := strings.Split(s, " ")

		score += getMove(battle[0], battle[1])

	}
	fmt.Println(score)
}
