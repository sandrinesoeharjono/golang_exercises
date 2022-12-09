package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func get_crab_positions() []int {
	// Create a slice of crab positions
	crab_positions := make([]int, 0)

	// Read the contents of the file
	file_contents, err := os.Open("crab_positions.txt")
	if err != nil {
		fmt.Println("There was an error when trying to read the file.")
		os.Exit(1)
	} else {
		scanner := bufio.NewScanner(file_contents)
		for scanner.Scan() {
			text := scanner.Text()
			// Split the numbers into a slice & transform to integers
			my_slice := strings.Split(text, ",")
			for _, s := range my_slice {
				n, _ := strconv.Atoi(s)
				crab_positions = append(crab_positions, n)
			}
		}
	}
	return crab_positions
}

func get_range_positions(my_slice []int) (int, int, int, int) {
	min_pos := 0
	max_pos := 0
	min_value := my_slice[0]
	max_value := my_slice[0]

	// Iterate over slice, finding the min/max values by comparing to previous min/max
	for i := 1; i < len(my_slice); i++ {
		if min_value > my_slice[i] {
			min_value = my_slice[i]
			min_pos = i
		}
		if max_value < my_slice[i] {
			max_value = my_slice[i]
			max_pos = i
		}
	}
	return min_value, min_pos, max_value, max_pos
}

func find_constant_moving_fuel(crab_pos []int, pos int) int {
	fuel_required := 0.0
	for n := 0; n < len(crab_pos); n++ {
		diff := pos - crab_pos[n]
		fuel_required += math.Abs(float64(diff))
	}
	return int(fuel_required)
}

func find_variable_moving_fuel(crab_pos []int, pos int) int {
	fuel_required := 0
	for n := 0; n < len(crab_pos); n++ {
		diff := pos - crab_pos[n]
		abs_diff := math.Abs(float64(diff))
		for i := 0; i < int(abs_diff); i++ {
			fuel_required += i + 1
		}
	}
	return int(fuel_required)
}

func main() {
	// Get crab positions from text file
	crab_pos := get_crab_positions()
	fmt.Printf("We have %d crabs.\n", len(crab_pos))

	// Determine min and max positions
	min_value, _, max_value, _ := get_range_positions(crab_pos)
	fmt.Printf("For crab positions: Min position = %d. Max position = %d.\n", min_value, max_value)

	// For all values from min-max, find the total fuel required to move the population (assuming constant fuel)
	constant_fuel_prices := make([]int, 0)
	for pos := min_value; pos < max_value; pos++ {
		fuel := find_constant_moving_fuel(crab_pos, pos)
		constant_fuel_prices = append(constant_fuel_prices, fuel)
	}

	// Find position with minimal fuel
	min_fuel, min_pos, max_fuel, max_pos := get_range_positions(constant_fuel_prices)
	fmt.Printf("For constant fuel: min fuel = %d (position %d). Max fuel = %d (position %d).\n", min_fuel, min_pos, max_fuel, max_pos)

	// For all values from min-max, find the total fuel required to move the population (assuming variable fuel)
	variable_fuel_prices := make([]int, 0)
	for pos := min_value; pos < max_value; pos++ {
		fuel := find_variable_moving_fuel(crab_pos, pos)
		variable_fuel_prices = append(variable_fuel_prices, fuel)
	}

	// Find position with minimal fuel
	min_fuel, min_pos, max_fuel, max_pos = get_range_positions(variable_fuel_prices)
	fmt.Printf("For variable fuel: min fuel = %d (position %d). Max fuel = %d (position %d).\n", min_fuel, min_pos, max_fuel, max_pos)
}
