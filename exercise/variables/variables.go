//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor string = "Green"
	fmt.Println("My favorite color is ", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	var birthYear, birthAge = 1990, 32
	fmt.Println("I was born in", birthYear, ". And I am", birthAge, "years old.")

	//* Store your first & last initials in two variables using block
	var (
		firstInitial string = "Y"
		lastInitial  string = "E"
	)
	fmt.Println("Initials = ", firstInitial, lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var birthDays int
	birthDays = 365 * birthAge
	fmt.Println("I am ", birthDays, "days old.")
}
