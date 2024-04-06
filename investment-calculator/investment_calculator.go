package main

import (
	"fmt"
	"math"
)

// it is common to import packages like this :)


func main() {
	var investmenrAmount float64 = 1000
	expectedReturnRate := 5.5 //Decleare or assign a variable where type should be infered 
	var years float64 = 10
	// variable declariton can be on the same line 
	// i.e var investmentAmount, years float64 = 1000, 10
	// i don't like that though I think its silly.

	
	//since go is hard static types all the variables in this need to be of the same type 
	var futureValue = investmenrAmount * math.Pow(1 + expectedReturnRate / 100, years)
	
	fmt.Println(futureValue)
}