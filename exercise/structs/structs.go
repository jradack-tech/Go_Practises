//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//--Requirements:
//* Create a rectangle structure containing its coordinates
type Rectangle struct {
	length int
	width  int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
func area(rect Rectangle) {
	fmt.Println("Area =", rect.length*rect.width)
}

func perimeter(rect Rectangle) {
	fmt.Println("Perimeter =", rect.length*2+rect.width*2)
}

func main() {
	rect := Rectangle{length: 10, width: 20}
	area(rect)
	perimeter(rect)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal
	rect.length = rect.length * 2
	rect.width = rect.width * 2
	area(rect)
	perimeter(rect)
}
