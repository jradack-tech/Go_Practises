//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type DirectVehicle interface {
	direct()
}

type Motorcycle string
type Car string
type Truck string

func (v Motorcycle) direct() {
	fmt.Println("- Motorcycle")
}

func (c Car) direct() {
	fmt.Println("- Car")
}

func (t Truck) direct() {
	fmt.Println("- Truck")
}

func direct(vehicles []DirectVehicle) {
	fmt.Println("All Vehicles are Ready ---")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		fmt.Printf("--Model: %v--\n", vehicle)
		vehicle.direct()
	}
	fmt.Println()
}

func main() {
	vehicles := []DirectVehicle{Motorcycle("Honda"), Car("Venz"), Truck("Isuzu")}
	direct(vehicles)
}
