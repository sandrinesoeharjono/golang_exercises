package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Test the program by running it and testing it by issuing three requests. Each request should involve a
different animal (cow, bird, snake) and a different property of the animal (eat, move, speak). Give 2
points for each request for which the program provides the correct response.

Examine the code. If the code contains a type called Animal, which is a struct containing three fields,
all of which are strings, then give another 2 points. If the program contains three methods called Eat(),
Move(), and Speak(), and all of their receiver types are Animal, give another 2 points.
*/

type animal struct {
	food       string
	locomotion string
	sound      string
}

func (a animal) Eat() string   { return a.food }
func (a animal) Move() string  { return a.locomotion }
func (a animal) Speak() string { return a.sound }

func main() {
	// Create a never-ending loop
	for true {
		// Define the prompt & initialize the input
		fmt.Printf(">")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')

		// Trim whitespace, set to lowercase & split into two variables
		user_string := strings.TrimSpace(input)
		lower_string := strings.ToLower(string(user_string))
		my_slice := strings.Split(lower_string, " ")
		requested_animal := my_slice[0]
		requested_info := my_slice[1]

		// Define the animal, if valid
		var my_animal animal
		if requested_animal == "cow" {
			my_animal.food = "grass"
			my_animal.locomotion = "walk"
			my_animal.sound = "moo"
		} else if requested_animal == "bird" {
			my_animal.food = "worms"
			my_animal.locomotion = "fly"
			my_animal.sound = "peep"
		} else if requested_animal == "snake" {
			my_animal.food = "mice"
			my_animal.locomotion = "slither"
			my_animal.sound = "hsss"
		} else {
			fmt.Println("This animal is not available.")
		}

		// Verify that info requested for that animal is valid
		if requested_info == "eat" {
			fmt.Println(my_animal.Eat())
		} else if requested_info == "move" {
			fmt.Println(my_animal.Move())
		} else if requested_info == "speak" {
			fmt.Println(my_animal.Speak())
		} else {
			fmt.Println("This information is not available.")
		}
	}
}
