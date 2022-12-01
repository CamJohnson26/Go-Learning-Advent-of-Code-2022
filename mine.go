package main

import (
	"fmt"
)

func main() {
	mystruct := struct {
		int32
		string
		testing string "This is a description of testing"
	}{32, "hello", "goodbye"}

	fmt.Print(mystruct)
	fmt.Print(mystruct.int32)
	fmt.Print(mystruct.testing)

	f := func(x, y int) int { return x + y }

	fmt.Print(f(100, 432))
	mystr := "testing"
	myarr := new([100]int)
	myslice := myarr[0:50]
	fmt.Println(myslice)
	//fmt.Print(len(mystr))

	for i, s := range mystr {
		fmt.Println(i, s)
	}
}
