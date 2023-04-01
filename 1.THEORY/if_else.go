package main

import "fmt"

// variable expression is available in if statement
func canIDring(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	result := canIDring(16)
	fmt.Println(result)
}
