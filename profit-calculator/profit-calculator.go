
// Profit Calculator
// CLI Tool in Go
// Ask User for Renvue, Expenses and Tax Rate
// Calualates Earning Before Tax and Earings After Tax and ratio of the two

package main

import (
	"fmt"
)

func main() {

	const taxRate float64 = 30
	var renvue float64
	var expenses float64

	fmt.Print("Please input your Revenue: ")
	fmt.Scan(&renvue)

	fmt.Print("Please input your Total Expenses: ")
	fmt.Scan(&expenses)

	preTaxRevenue := renvue - expenses
	var taxFloat float64 = 1 - taxRate / 100

	postTaxRevenue := preTaxRevenue * taxFloat

	ratio := preTaxRevenue / postTaxRevenue

	fmt.Print("Yor pre tax revenue is: ")
	fmt.Println(preTaxRevenue)

	fmt.Print("Your post tax revenue is: ")
	fmt.Println(postTaxRevenue)

	fmt.Print("The Ratio is: ")
	fmt.Println(ratio)


	



}