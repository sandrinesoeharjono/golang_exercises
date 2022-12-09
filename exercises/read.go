package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file
has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname
for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found
in the file. Each struct created will be added to a slice, and after all lines have been read from
the file, your program will have a slice containing one struct for each line in the file. After
reading all lines from the file, your program should iterate through your slice of structs and
print the first and last names found in each struct.
*/

func main() {
	// Create a 'Name' structure with 2 fields: fname & lname (each a string of 20 chars)
	type Name struct {
		fname string
		lname string
	}

	// Declare a slice of the Name structs (size 0 initially)
	my_slice := make([]Name, 0)

	// Prompt the user for a file name
	var filename string
	fmt.Printf("Please enter a file name: ")
	fmt.Scan(&filename)

	// Read the contents of the file
	file_contents, err := os.Open(filename)
	if err == nil {
		// Iterate over each line in the file's contents
		scanner := bufio.NewScanner(file_contents)
		for scanner.Scan() {
			// Split first and last names
			words := strings.Split(scanner.Text(), " ")

			// Create a struct from the data
			one_person := Name{words[0], words[1]}

			// Add to the slice
			my_slice = append(my_slice, one_person)
		}
	} else {
		fmt.Println("There was an error when trying to read the file:", filename)
	}

	// Iterate over each item in the slice to print each name
	for i := 0; i < len(my_slice); i++ {
		fmt.Println(my_slice[i].fname, my_slice[i].lname)
	}

	// Close file
	file_contents.Close()
}
