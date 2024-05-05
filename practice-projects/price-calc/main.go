package main

import "fmt"

func main() {
	prices := []float64{10 ,20 ,30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIcluded []float64 = make([]float64, len(prices))
		for x, price := range prices {
			taxIcluded[x] = price * (1+taxRate)
		}
		result[taxRate] = taxIcluded
	}

	fmt.Println(result)
}