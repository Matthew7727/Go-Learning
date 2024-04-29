package main

import (
	"fmt"
)

func main() {

	prices := []float64{9.99, 7.50}
	fmt.Println(prices, len(prices))
	prices = append(prices, 5.99)
	fmt.Println(prices, len(prices))
}