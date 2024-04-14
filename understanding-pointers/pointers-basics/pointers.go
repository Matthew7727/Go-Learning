package main

import "fmt"

func main() {
	age := 32 // regular varibale
	agePointer := &age

	fmt.Println("age: ", *agePointer)
	fmt.Println("adult years: ", getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}