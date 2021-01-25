package main

import (
	"fmt"
)

func main() {
	array := []string{"h", "e", "l", "l", "o"}
	fmt.Printf("%v", reverse(array))
}

func reverse(array []string) []string {
	result := make([]string, 0, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		result = append(result, array[i])
	}
	return result
}
