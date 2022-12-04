package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isOverlap(a1, a2, b1, b2 int) bool {
	return (a2 >= b1 && a1 <= b1) || (b2 >= a1 && b1 <= a1)
}

func Part2() {
	//myFile, _ := os.ReadFile("221204/test_input.txt")
	myFile, _ := os.ReadFile("221204/input.txt")

	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")
	score := 0
	for _, line := range lines {
		vals := strings.Split(line, ",")
		a := vals[0]
		b := vals[1]
		a_vals := strings.Split(a, "-")
		b_vals := strings.Split(b, "-")
		a1, _ := strconv.Atoi(a_vals[0])
		a2, _ := strconv.Atoi(a_vals[1])
		b1, _ := strconv.Atoi(b_vals[0])
		b2, _ := strconv.Atoi(b_vals[1])
		fmt.Println(a1, a2, b1, b2)
		contained := isOverlap(a1, a2, b1, b2)
		if contained {
			score += 1
		}
	}
	fmt.Println(score)
}
