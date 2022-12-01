package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
	myFile, _ := os.ReadFile("221201/input.txt")
	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")

	highest := 0
	highest_1 := 0
	highest_2 := 0
	cur := 0
	for _, s := range lines {
		if s == "" {
			if cur > highest {
				highest_2 = highest_1
				highest_1 = highest
				highest = cur
			} else
			if cur > highest_1 {
				highest_2 = highest_1
				highest_1 = cur
			} else
			if cur > highest_2 {
				highest_2 = cur
			}
			fmt.Println(highest, highest_1, highest_2)
			cur = 0
		} else {
			sval, _ := strconv.Atoi(s)
			cur += sval
		}
	}
	if cur > highest {
		highest = cur
	}
	fmt.Println(highest + highest_1 + highest_2)
}
