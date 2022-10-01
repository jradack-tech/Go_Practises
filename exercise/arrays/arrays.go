//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
type Product struct {
	name  string
	price int
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	var list [4]Product

	//* Insert 3 products into the array
	list[0] = Product{name: "clock", price: 100}
	list[1] = Product{name: "computer", price: 500}
	list[2] = Product{name: "book", price: 5}

	//* Print to the terminal:
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	fmt.Println("The last item on the list is", list[len(list)-1].name)
	fmt.Println("The total number of items is", len(list))

	var totalCost int
	for i := 0; i < len(list); i++ {
		totalCost += list[i].price
	}
	fmt.Println("The total cost of the items is $", totalCost)

	//* Add a fourth product to the list and print out the
	//  information again
	fmt.Println("Adding a fourth product now...")
	list[3] = Product{name: "Dress", price: 50}

	fmt.Println("The last item on the list is", list[len(list)-1].name)
	fmt.Println("The total number of items is", len(list))

	totalCost = 0
	for i := 0; i < len(list); i++ {
		totalCost += list[i].price
	}
	fmt.Println("The total cost of the items is $", totalCost)
}
