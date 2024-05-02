package main

import "fmt"

func main() {

	websites := map[string]string{
		"Go":  "https://google.com",
		"AWS": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["AWS"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	websites["Go"] = "https://go.com"
	delete(websites, "AWS")
	fmt.Print(websites)



}