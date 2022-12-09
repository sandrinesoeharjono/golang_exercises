package main

import "fmt"

/*
Assignment: Write a program which prompts the user to enter a floating point number and
prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/

func main() {
	// Define the question
	fmt.Printf("Please enter a floating point number.")

	// Initialize the float input
	var float_input float64

	// Use format to ask the user's input
	fmt.Scan(&float_input) // we ignore the number of inputs / error messages here

	// Transform to integer (i.e truncate)
	var int_input int = int(float_input)

	// Print the truncated result
	fmt.Println(int_input)
}
