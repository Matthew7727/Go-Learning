package main

import  "fmt"

func main() {

	fmt.Println("Welcome to MattBank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money?")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	var balance float64 = 10000

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	
	if (choice == 1) {
		fmt.Printf("Your balance is: £%.2f \n", balance)
	} else if (choice == 2) {
		var money int 
		fmt.Print("Your deposit: £")
		fmt.Scan(&money)
		balance += float64(money)
		fmt.Printf("Your new balance is: £%.2f \n", balance)
	} else if (choice == 3) {
		var money int
		fmt.Print("Your withdrawal: £")
		fmt.Scan(&money)
		balance -= float64(money)
		fmt.Printf("Your new balance is: £%.2f \n", balance)
	} else {
		fmt.Println("Goodbye")
	}




}