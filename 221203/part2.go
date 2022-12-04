package main

import (
	"fmt"
	"os"
	"strings"
)

func getBadge(groups []string) string {

	found_map := map[string]int{}

	for i, group := range groups {
		found_map_group := map[string]bool{}
		for _, c := range group {
			char := string(c)
			if found_map_group[char] == false {
				found_map[char] += i + 1
				found_map_group[char] = true
			}
		}
	}

	for k, v := range found_map {
		if v == 6 {
			return k
		}
	}
	return "error"
}

func Part2() {
	//myFile, _ := os.ReadFile("221203/test_input.txt")
	myFile, _ := os.ReadFile("221203/input.txt")

	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")
	score := 0

	groups := len(lines) / 3
	for groupIndex := 0; groupIndex < groups; groupIndex++ {
		group := lines[groupIndex*3 : groupIndex*3+3]
		
		item := getBadge(group)
		if item == "error" {
			fmt.Println("error")
			break
		}
		score += GetPriority(item)
		//fmt.Println(getPriority(item))
		//score += getPriority(item)
	}
	fmt.Println(score)
}
