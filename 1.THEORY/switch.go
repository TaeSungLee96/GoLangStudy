package main

import "fmt"

// variable expression is available in switch statement
func canIDring(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	result := canIDring(16)
	fmt.Println(result)
}
