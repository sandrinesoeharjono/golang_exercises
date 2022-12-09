package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON
object from the map, and then your program should print the JSON object.
*/

func main() {
	// Create a reader
	inputReader := bufio.NewReader(os.Stdin)

	// Prompt the user for a name
	fmt.Printf("Please enter a name: ")
	input, _ := inputReader.ReadString('\n')
	name_input := strings.TrimSpace(input)

	// Prompt the user for an address
	fmt.Printf("Please enter an address: ")
	add_input, _ := inputReader.ReadString('\n')
	address_input := strings.TrimSpace(add_input)

	// Create a map and address the name/address to it
	my_obj := map[string]string{"name": name_input, "address": address_input}

	// Create a JSON object
	json_object, err := json.Marshal(my_obj)

	if err == nil {
		// Print the JSON object
		fmt.Println("The JSON rendering is the following:", string(json_object))
	} else {
		fmt.Println("The provided inputs are not JSON-compatible.")
	}
}
