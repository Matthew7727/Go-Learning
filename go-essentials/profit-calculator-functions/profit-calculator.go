
// Profit Calculator
// CLI Tool in Go
// Ask User for Renvue, Expenses and Tax Rate
// Calualates Earning Before Tax and Earings After Tax and ratio of the two

package main

import (
	"fmt"
)

const taxRate float64 = 30

func main() {

	revenue := outputText("Please input your Revenue: ")
	expenses := outputText("Please input your Total Expenses: ")

	preTaxProfit, postTaxProfit := calculateProfits(revenue, expenses)

	ratio := calcuateRatio(preTaxProfit, postTaxProfit)

	fmt.Printf("Yor pre tax revenue is: £%.2f \n ", preTaxProfit)

	fmt.Printf("Your post tax revenue is: £%.2f \n ", postTaxProfit)

	fmt.Printf("The Ratio is: %.2f \n", ratio)

}

func outputText(text string) (input float64) {
	fmt.Print(text)
	fmt.Scan(&input)
	return
}

func calculateProfits(revenue float64, expenses float64) (float64, float64) {
	preTaxRevenue := revenue - expenses
	var taxFloat float64 = 1 - taxRate / 100

	postTaxRevenue := preTaxRevenue * taxFloat

	return preTaxRevenue, postTaxRevenue
}

func calcuateRatio(preTax, postTax float64) (float64) {
	return preTax / postTax
}