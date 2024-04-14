package main

import ( 
	"fmt"
	"os"
	"strconv"
)

const FILE_NAME string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(FILE_NAME, []byte(balanceTxt), 0644)
}

func readBalanceFromFile() (float64) {
	data, _ := os.ReadFile(FILE_NAME)
	balanceTxt := string(data)
	floatBalance, _ := strconv.ParseFloat(balanceTxt, 64)
	
	return floatBalance 
}

func main() {
	for {
		fmt.Println("Welcome to MattBank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money?")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		var balance float64 = readBalanceFromFile()

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
			writeBalanceToFile(balance)
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
			writeBalanceToFile(balance)
			fmt.Printf("Your new balance is: £%.2f \n", balance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thank for choosing our bank!")
			return
		}
	}
}
