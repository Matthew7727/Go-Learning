package main

import (
	"fmt"
	"math"
)

// it is common to import packages like this :)

func main() {

	const inflationRate = 3.8

	var investmenrAmount float64 
	var expectedReturnRate float64
	var years float64

	outputText("Pleas input your investment amount: ")
	fmt.Scan(&investmenrAmount) //  this is pointer we'll come back to these

	fmt.Print("Please input your expected retrun rate: ")
	fmt.Scan(&expectedReturnRate) // fmt.Scan has limitaions

	fmt.Print("How long will this investment be for (in years): ")
	fmt.Scan(&years)

	
	futureValue := investmenrAmount * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFVOutput := fmt.Sprintf("Future Value: $%.2f\n", futureValue)

	formattedAFIOutput :=  fmt.Sprintf("Future Value (adjusted for Inflation): $%.2f\n", inflationAdjustedValue)
	fmt.Println(formattedFVOutput, formattedAFIOutput)
}

func outputText(text string) {
	fmt.Print(text)
}