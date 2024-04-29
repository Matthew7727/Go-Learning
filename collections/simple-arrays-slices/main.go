package main

import "fmt"


func main() {
	var productName [4]string 
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	productName[3] = "Hello Darkness "

	fmt.Println(prices)
	fmt.Println(productName)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	fmt.Println(len(productName), cap(featuredPrices))

}

