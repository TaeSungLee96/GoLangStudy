package main

import (
	"fmt"
	"strings"
)

func multiply1(a int, b int) int {
	return a * b
}

// a, b are of type int
func multiply2(a, b int) int {
	return a * b
}

// lenAndUpper returns length of string and upper case of string
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// "...string" means that we can pass any number of strings
func reaptMe(words ...string) {
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
}

func main() {
	fmt.Println(multiply1(2, 3))
	fmt.Println(multiply2(2, 3))

	totalLen, upperName := lenAndUpper("TAESUNG")
	fmt.Println(totalLen, upperName)
	// or We can ignore the value we don't need with "_"
	// totalLen, _ := lenAndUpper("TAESUNG")

	reaptMe("taesung", "taemin", "taehyung")
}
