// communication.go

package main

import "fmt"
import "github.com/pallinder/go-randomdata"

func presentOptions() {
	fmt.Println("Welcome to MattBank")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money?")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}