package main

import "fmt"

func main() {
	age := 32 // regular varibale
	agePointer := &age

	fmt.Println("age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}