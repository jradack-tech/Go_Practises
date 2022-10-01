//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//

package main

import "fmt"

//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
type Product struct {
	name     string
	security bool
}

//* Create functions to activate and deactivate security tags using pointers
func changeSecurity(product *Product, state bool) {
	product.security = state
	if state {
		fmt.Println("The security of", product.name, "is activated")
	} else {
		fmt.Println("The security of", product.name, "is deactivated")
	}
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(products []Product) {
	for _, product := range products {
		changeSecurity(&product, false)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	products := []Product{
		{name: "Product 1"},
		{name: "Product 2"},
		{name: "Product 3"},
		{name: "Product 4"},
	}

	//  - Deactivate any one security tag in the array/slice
	fmt.Println("Deactivating any one security tag...")
	changeSecurity(&products[0], false)

	//  - Call the checkout() function to deactivate all tags
	fmt.Println("Deactivating all tags...")
	checkout(products)

	//  - Print out the array/slice after each change
	for _, product := range products {
		fmt.Println("The security of", product.name, "is activate state:", product.security)
	}
}
