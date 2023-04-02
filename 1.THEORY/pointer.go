package main

import "fmt"

func main() {
	a := 2
	b := &a // b is a pointer to a
	*b = 20 // a is changed to 20
	fmt.Println(a) // 20
}
