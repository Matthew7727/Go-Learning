// Profit Calculator
// CLI Tool in Go
// Ask User for Renvue, Expenses and Tax Rate
// Calualates Earning Before Tax and Earings After Tax and ratio of the two
// Validates User Input
// - No Negative Numbers
// - Not 0
// Stores Resutls into a file

package main

import (
	"errors"
	"fmt"
	"os"
)

const taxRate float64 = 30

func main() {

	revenue, RevErr := getUserInput("Please input your Revenue: ")
	if RevErr != nil {
		panic("Error with revenue input")
	}

	expenses, ExpErr := getUserInput("Please input your Total Expenses: ")
	if ExpErr != nil {
		panic("Error with expenese input")
	}
	preTaxProfit, postTaxProfit := calculateProfits(revenue, expenses)

	ratio := calcuateRatio(preTaxProfit, postTaxProfit)

	fmt.Printf("Yor pre tax revenue is: £%.2f \n ", preTaxProfit)
	fmt.Printf("Your post tax revenue is: £%.2f \n ", postTaxProfit)
	fmt.Printf("The Ratio is: %.2f \n", ratio)

	saveToFile(preTaxProfit, postTaxProfit, ratio)

}

func saveToFile(preTax float64, postTax float64, ratio float64) {
	stringToWrite := fmt.Sprintf("The Pre Tax Proft was: £%.2f. \n The Post Tax Proft was: £%.2f \n with a ratio of %.2f", preTax, postTax, ratio)
	os.WriteFile("Profts.txt", []byte(stringToWrite), 0644)
}

func getUserInput(text string) (float64, error) {
	var input float64

	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("you have entered an invalid amount")
	} else {
		return input, nil
	}
}

func calculateProfits(revenue float64, expenses float64) (float64, float64) {
	preTaxRevenue := revenue - expenses
	var taxFloat float64 = 1 - taxRate/100

	postTaxRevenue := preTaxRevenue * taxFloat

	return preTaxRevenue, postTaxRevenue
}

func calcuateRatio(preTax, postTax float64) float64 {
	return preTax / postTax
}
