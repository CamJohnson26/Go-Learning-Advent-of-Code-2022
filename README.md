# Go-Learning-Advent-of-Code-2022

Using 2022 Advent of Code to play around with Go. Advent of code: https://adventofcode.com/2022

## Links
Best resource so far is the spec: https://go.dev/ref/spec#Array_types

And also the docs for standard libs:
https://pkg.go.dev/std
https://pkg.go.dev/builtin

Cheatsheet:
https://github.com/a8m/golang-cheat-sheet

Best IDE: Using GoLand: https://www.jetbrains.com/go/

## Some common operations:

Iterate lines of a file:
```
myFile, _ := os.ReadFile("221204/input.txt")

	myStr := string(myFile)
	lines := strings.Split(myStr, "\n")
	
	for _, line := range lines {
    fmt.Println(line)
  }
 ```
 
 Cast a string to an int:
 ```
 myInt, _ := strconv.Atoi(myStr)
 ```
 
 Create a map:
 ```
 myMap := map[string]bool{}
 myMap["key"] = false
 ```
