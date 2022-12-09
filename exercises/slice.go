package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop, the program prompts the user to
enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any
number of integers which the user decides to enter. The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer.
*/

func main() {
	// Create an empty integer slice of length 3
	my_slice := make([]int, 3)

	for i := 0; i >= 0; i++ {
		// Prompt the user for an integer input
		fmt.Print("Please enter an integer: ")
		inputReader := bufio.NewReader(os.Stdin)
		value, _ := inputReader.ReadString('\n')
		value = strings.TrimSpace(value) // Trim whitespace

		if value == "X" {
			// Only input that can exit the loop
			break
		} else {
			// Attempt to convert string value to an integer
			int_value, error := strconv.Atoi(string(value))
			if error == nil {
				// If conversion was successful, append to slice
				my_slice = append(my_slice, int_value)

				// Sort slice and print it out
				sort.Ints(my_slice)
				fmt.Println("Current slice results:", my_slice)
			} else {
				// If conversion was unsuccessful, print an informative message and continue to the next loop
				fmt.Println("The input", value, "is not a valid integer - Please try again.")
				continue
			}
		}
	}

	// Print the final result
	fmt.Println("Final slice result:", my_slice)
}
