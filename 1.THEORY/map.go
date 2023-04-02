package main

import "fmt"

// map is similar to object in JavaScript
// map[keyType]valueType
func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
}
