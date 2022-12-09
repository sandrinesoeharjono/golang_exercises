package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite
search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function BubbleSort() which takes a slice of integers as an argument
and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent
elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should
take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap()
function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

Test the program by running it twice and testing it with a different sequence of integers each time. The first test
sequence of integers should be all positive numbers and the second test should have at least one negative number.
Give 3 points if the program works correctly for one test sequence, and give 3 more points if the program works
correctly for the second test sequence.

Examine the code. If the code contains a function called BubbleSort() which takes a slice of integers as an argument,
then give another 2 points. If the code contains a function called Swap() function which takes two arguments, a
slice of integers and an index value i, then give another 2 points.
*/

func Swap(my_slice []int, i int) {
	// Get value at indexes i & i+1
	value_i := my_slice[i]
	value_after := my_slice[i+1]

	// Swap values if not in ascending order
	if value_i > value_after {
		my_slice[i] = value_after
		my_slice[i+1] = value_i
	}
}

func BubbleSort(my_slice []int) {
	for n_iterations := 0; n_iterations < 10; n_iterations++ {
		for i := 0; i < len(my_slice)-1; i++ {
			Swap(my_slice, i)
		}
	}
}

func main() {
	// Prompt user for input
	fmt.Print("Please enter a comma-separated sequence (of maximum length 10): ")
	inputReader := bufio.NewReader(os.Stdin)
	value, _ := inputReader.ReadString('\n')
	value = strings.TrimSpace(value) // Trim whitespace

	// Split into a slice
	my_slice := strings.Split(value, ",")

	// Ensure that the slice is not longer than 10
	if len(my_slice) > 10 {
		fmt.Println("The sequence cannot be of length longer than 10.")
		os.Exit(1)
	}

	// Convert strings to int
	int_slice := make([]int, len(my_slice))
	for i, s := range my_slice {
		int_slice[i], _ = strconv.Atoi(s)
	}
	fmt.Println("You inputted:", int_slice)

	// Call 'BubbleSort' to sort the slice
	BubbleSort(int_slice)

	// Print output (1 item per line)
	fmt.Println("The sorted results are the following:")
	for i := 0; i < len(int_slice); i++ {
		fmt.Println(int_slice[i])
	}
}
