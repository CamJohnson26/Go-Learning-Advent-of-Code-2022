package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myFile, _ := os.ReadFile("221201/input.txt")
	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")
	fmt.Println(lines)

	highest := 0
	cur := 0
	for _, s := range lines {
		if s == "" {
			if cur > highest {
				highest = cur
			}
			cur = 0
		} else {
			sval, _ := strconv.Atoi(s)
			cur += sval
		}
	}
	if cur > highest {
		highest = cur
	}
	fmt.Println(highest)
	part2()

}
