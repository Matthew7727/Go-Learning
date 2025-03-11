package main

import (
	//"fmt"
	"fmt"
	cmdmanager "exapmle.com/practice-cacl/cmdmanager"
	// filemanager "exapmle.com/practice-cacl/file-manager"
	prices "exapmle.com/practice-cacl/prices"
)



func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json",taxRate*100))
		cmdmanager := cmdmanager.New()
		priceJob := prices.NewTaxInclPriceJob(cmdmanager, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println(err)
			fmt.Println("something happened whoops")
		}
	}
}