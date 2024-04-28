package main

import (
	"fmt"
)

func main() {

	result := add(5, 7)
	fmt.Println(result)

	strResult := add("Hello ", "World ")
	fmt.Println(strResult)

}

func add[T int|float64|string ](a, b T) T {

	return a + b
}