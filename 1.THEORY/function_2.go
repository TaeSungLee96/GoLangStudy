package main

import (
	"fmt"
	"strings"
)

// This function "nasked" return
// You don't need to write return "value"
func lenAndUpper(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func main() {
	totalLen, upperName := lenAndUpper("TAESUNG")
	fmt.Println(totalLen, upperName)
}
