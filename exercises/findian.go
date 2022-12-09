package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to enter a string. The program searches through the
entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!”
if the entered string starts with the character ‘i’, ends with the character ‘n’, and
contains the character ‘a’. The program should print “Not Found!” otherwise. The program
should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings,
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!”
for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

func main() {
	// Define the prompt
	fmt.Printf("Please enter a string: ")

	// Ask the user's input (use bufio NewReader to accept a multi-word string)
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	// Trim whitespace & set string to lowercase
	user_string := strings.TrimSpace(input)
	lower_string := strings.ToLower(string(user_string))

	// Check the 3 conditions: presence of 'i',' 'a' and 'n' at certain indexes
	starts_i := strings.HasPrefix(lower_string, "i")  // return True if string starts with 'i'
	ends_n := strings.HasSuffix(lower_string, "n")    // return True if the string ends with 'n'
	contains_a := strings.Contains(lower_string, "a") // return True if string contains 'a'

	// Check the status of all 3 conditions
	if starts_i && ends_n && contains_a {
		// Print "Found!" if all 3 conditions evaluate to 3
		fmt.Println("Found!")
	} else {
		// Otherwise print 'Not Found!'
		fmt.Println("Not Found!")
	}
}
