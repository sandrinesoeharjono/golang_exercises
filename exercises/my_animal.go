package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

// Define the 3 animals: cow, bird & snake
type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

// Define the animal-related functions Eat(), Move(), Speak() and getName()
func (c Cow) Eat()              { fmt.Println("grass") }
func (c Cow) Move()             { fmt.Println("walk") }
func (c Cow) Speak()            { fmt.Println("moo") }
func (c Cow) getName() string   { return c.name }
func (b Bird) Eat()             { fmt.Println("worms") }
func (b Bird) Move()            { fmt.Println("fly") }
func (b Bird) Speak()           { fmt.Println("peep") }
func (b Bird) getName() string  { return b.name }
func (s Snake) Eat()            { fmt.Println("mice") }
func (s Snake) Move()           { fmt.Println("slither") }
func (s Snake) Speak()          { fmt.Println("hsss") }
func (s Snake) getName() string { return s.name }

func main() {
	// Create an empty slice of animals
	var my_animals []Animal

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

		// Identify type of request (newanimal or query)
		if len(my_slice) == 3 {
			request_type := my_slice[0]
			switch request_type {
			case "newanimal":
				// Create a new animal
				animal_name := my_slice[1]
				animal_type := my_slice[2]
				if animal_type == "bird" {
					my_animals = append(my_animals, Bird{name: animal_name})
					fmt.Println("Created the bird!")
				} else if animal_type == "cow" {
					my_animals = append(my_animals, Cow{name: animal_name})
					fmt.Println("Created the cow!")
				} else if animal_type == "snake" {
					my_animals = append(my_animals, Snake{name: animal_name})
					fmt.Println("Created the snake!")
				} else {
					fmt.Println("Invalid animal", animal_type)
				}

			case "query":
				// Identify the requested info for an existing animal
				animal_name := my_slice[1]
				info_requested := my_slice[2]
				for _, animal := range my_animals {
					if animal.getName() == animal_name {
						switch info_requested {
						case "move":
							animal.Move()
						case "eat":
							animal.Eat()
						case "speak":
							animal.Speak()
						default:
							fmt.Println("This info isn't available: ", info_requested)
						}
					} else {
						fmt.Println("This animal does not exist.")
					}
				}

			default:
				fmt.Println("Request cannot be processed.")
			}
		} else {
			fmt.Println("Request must contain 3 words.")
		}
	}
}
