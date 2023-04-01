package main

import "fmt"

func multiply1(a int, b int) int {
	return a * b
}
func multiply2(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply1(2, 3))
	fmt.Println(multiply2(2, 3))
}
