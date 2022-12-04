package main

import (
	"fmt"
	"os"
	"strings"
)

func getCommonItem(input string) string {
	length := len(input)
	midpoint := length / 2
	a := input[:midpoint]
	b := input[midpoint:]
	fmt.Println(a, b)
	found_map := map[string]bool{}
	dup_char := ""

	for _, c := range a {
		char := string(c)
		found_map[char] = true
	}
	for _, c := range b {
		char := string(c)
		found := found_map[char]
		if found {
			dup_char = char
		}
	}

	return dup_char
}

func GetPriority(item string) int {
	index := int(rune(item[0]))
	if index > 65 && index <= 96 {
		return index - 38
	} else if index > 96 && index <= 123 {
		return index - 96
	}
	return 0
}

func Part1() {
	//myFile, _ := os.ReadFile("221203/test_input.txt")
	myFile, _ := os.ReadFile("221203/input.txt")

	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")
	score := 0
	for _, line := range lines {
		item := getCommonItem(line)
		fmt.Println(item)
		fmt.Println(GetPriority(item))
		score += GetPriority(item)
	}
	fmt.Println(score)
}
