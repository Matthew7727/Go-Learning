package main

import (
	"exapmle.com/practice-cacl/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxInclPriceJob(taxRate)
		priceJob.Process()
	}
}