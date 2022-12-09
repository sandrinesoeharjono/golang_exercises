package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/* Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the
array should print the subarray that it will sort. When sorting is complete, the main goroutine should
print the entire sorted list. */

func sort_array(wg *sync.WaitGroup, array []int) []int {
	defer wg.Done()
	fmt.Println("=> Goroutine sorts the array:", array)
	// Sorts an array in place
	sort.Ints(array)
	return array
}

func merge_arrays(array1 []int, array2 []int, array3 []int, array4 []int) []int {
	// Merges 4 arrays together
	new_array := []int{}
	new_array = append(array1, array2...)
	new_array = append(new_array, array3...)
	new_array = append(new_array, array4...)
	fmt.Println("Merged arrays:", new_array)
	return new_array
}

func main() {
	fmt.Println("Please input a series of integers, separated by commas:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	// Trim whitespace, set to lowercase & split into two variables
	user_string := strings.TrimSpace(input)
	lower_string := strings.ToLower(string(user_string))
	string_array := strings.Split(lower_string, ",")

	// Convert []string to []int
	my_array := make([]int, 0)
	for _, s := range string_array {
		n, _ := strconv.Atoi(s)
		my_array = append(my_array, n)
	}

	// Separate the array into 4 parts (slices)
	array_size := len(my_array) / 4
	slice1 := my_array[:array_size]
	slice2 := my_array[array_size : 2*(array_size)]
	slice3 := my_array[2*(array_size) : 3*(array_size)]
	slice4 := my_array[3*(array_size):]

	// Create a waitgroup with 4 concurrent goroutines (1 per slice)
	// Each goroutine will sort a slice of integers
	var wg sync.WaitGroup
	wg.Add(4)
	go sort_array(&wg, slice1)
	go sort_array(&wg, slice2)
	go sort_array(&wg, slice3)
	go sort_array(&wg, slice4)
	wg.Wait() // ensure all goroutines completed before proceeding

	// Merge all 4 slices together
	final_array := merge_arrays(slice1, slice2, slice3, slice4)

	// Run final sorting of final array
	sort.Ints(final_array)
	fmt.Println("Final sorted result:", final_array)
}
