package main

import "fmt"


// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func main() {

	// 1	
	hobbies := [3]string{"Coding", "Apex", "Gym"}
	
	// 2
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3
	hobbiesSlice1 := hobbies[:2]
	hobbiesSlice2 := hobbies[0:2]

	fmt.Println(hobbiesSlice1, hobbiesSlice2)

	// 4 
	hobbiesSlice1 = hobbiesSlice1[1:3]
	fmt.Println(hobbiesSlice1)

	// 5
	var goals = []string{"Get a job at Monzo", "Learn Go"}
	fmt.Println(goals)
	
	// 6 
	goals[1] = "Get Really Good at Go"
	goals = append(goals, "Move to the US")
	fmt.Println(goals)

	// 7 
	type product struct {
		title string
		id int64
		price float64
	}

	prod1 := product{
		title: "Phone",
		id: 1,
		price: 999.99,
	}

	prod2 := product{
		title: "Tablet",
		id: 2,
		price: 1999.99,
	}

	prod3 := product{
		title: "Laptop",
		id: 3,
		price: 1299.99,
	}
	

	products := []product{prod1, prod2}
	fmt.Println(products)
	products = append(products, prod3)
	fmt.Println(products)

}

