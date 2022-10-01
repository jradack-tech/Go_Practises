//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(personName string) {
	fmt.Println("Hello, My name is", personName)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func returnAnyMessage() string {
	return "I love Golang!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func sum(a int, b int, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func returnAnyNumber(num int) int {
	return num
}

//* Write a function that returns any two numbers
func returnAnyTwoNumbers(num1 int, num2 int) (int, int) {
	return num1, num2
}

func main() {
	greeting("Eugene")

	fmt.Println("Return any message:", returnAnyMessage())

	fmt.Println("5 + 6 + 7 =", sum(5, 6, 7))

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	//* Call every function at least once
	num1, num2 := returnAnyTwoNumbers(5, 6)
	fmt.Println("5 + 6 + 8 =", sum(num1, num2, returnAnyNumber(8)))
}
