package main

import (
	"fmt"
	"strings"
)

// defer is used to execute a function after the function is done
func lenAndUpper(name string) (length int, upperName string) {
	defer fmt.Println("I'm done")
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func main() {
	totalLen, upperName := lenAndUpper("TAESUNG")
	fmt.Println(totalLen, upperName)
}
