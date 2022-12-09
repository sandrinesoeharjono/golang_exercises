package main

import (
	"fmt"
	"math"
)

/*
Write a program which first prompts the user to enter values for acceleration, initial velocity,
and initial displacement. Then the program should prompt the user to enter a value for time
and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time, assuming the given
values acceleration, initial velocity, and initial displacement. The function returned by
GenDisplaceFn() should take one float64 argument t, representing time, and return one float64
argument which is the displacement travelled after time t
*/

func request_user_inputs() (float64, float64, float64, float64) {
	// Initialize the inputs
	var acceleration float64
	var velocity float64
	var displacement float64
	var time float64
	var user_input float64

	// Ask for acceleration input
	fmt.Printf("Please enter a value for the acceleration: ")
	fmt.Scan(&user_input)
	acceleration = float64(user_input)

	// Ask for initial velocity input
	fmt.Printf("Please enter a value for the initial velocity: ")
	fmt.Scan(&user_input)
	velocity = float64(user_input)

	// Ask for initial displacement input
	fmt.Printf("Please enter a value for the initial displacement: ")
	fmt.Scan(&user_input)
	displacement = float64(user_input)

	// Ask for time input
	fmt.Printf("Please enter a value for the time: ")
	fmt.Scan(&user_input)
	time = float64(user_input)

	return acceleration, velocity, displacement, time
}

func GenDisplaceFn(a float64, vo float64, displacement float64) func(t float64) float64 {
	return func(t float64) float64 { return 0.5*a*math.Pow(t, 2) + vo*t + displacement }
}

func main() {
	// Get user inputs
	acc, velocity, displacement, time := request_user_inputs()
	fmt.Printf("Values provided: acceleration: %f, velocity: %f, displacement: %f, time %f \n", acc, velocity, displacement, time)

	// Create function that depends strictly on time
	dist_function := GenDisplaceFn(acc, velocity, displacement)
	distance := dist_function(time)
	fmt.Println("Distance traveled:", distance)
}
