package main

import (
	"fmt"
	"example.com/bank-app/fileops"

)

const FILE_NAME string = "balance.txt"

func main() {
	for {
		presentOptions()

		var choice int
		balance , err  := fileops.GetFloatFromFile(FILE_NAME)

		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("===============")
			panic("Ended in a panic")
		}

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		

		switch choice {
		case 1:
			fmt.Printf("Your balance is: £%.2f \n", balance)
		case 2:
			var money int
			fmt.Print("Your deposit: £")
			fmt.Scan(&money)
			if money <= 0 {
				fmt.Println("This is an invlaid amount")
				continue // this doens't just stop this else statment it stops all code from being exexcuted from that point
			}
			balance += float64(money)
			fileops.WriteFloatToFile(balance, FILE_NAME)
			fmt.Printf("Your new balance is: £%.2f \n", balance)
		case 3:
			var money int
			fmt.Print("Your withdrawal: £")
			fmt.Scan(&money)
			if money <= 0 {
				fmt.Println("This is an invlaid amount")
				continue
			}
			if float64(money) > balance {
				fmt.Println("You cannot withdraw more money than u have silly goose")
				continue
			}
			balance -= float64(money)
			fileops.WriteFloatToFile(balance, FILE_NAME)
			fmt.Printf("Your new balance is: £%.2f \n", balance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thank for choosing our bank!")
			return
		}
	}
}


