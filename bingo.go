package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fill_initial_data() ([]string, [][][]string) {
	// Create a slice of bingo numbers
	bingo_numbers := make([]string, 0)

	// Create a slice of boards (aka list of 2D arrays)
	boards := make([][][]string, 0)

	// Read the contents of the file
	file_contents, err := os.Open("octopus.txt")
	if err != nil {
		fmt.Println("There was an error when trying to read the file.")
		os.Exit(1)
	} else {
		// Iterate over each line in the file's contents
		scanner := bufio.NewScanner(file_contents)
		line_number := 0

		// Create an initial empty board; we'll fill its entries one by one
		my_board := make([][]string, 5)
		board_entry := 0

		for scanner.Scan() {
			text := scanner.Text()
			if line_number == 0 {
				// Capture selected bingo numbers
				my_slice := strings.Split(text, ",")
				for _, s := range my_slice {
					//bingo_numbers[i], _ = strconv.Atoi(s)
					//n, _ := strconv.Atoi(s)
					bingo_numbers = append(bingo_numbers, s)
				}
				fmt.Println("Numbers:", bingo_numbers, "\n")
			} else if line_number > 1 {
				if text == "" {
					// End of existing board
					boards = append(boards, my_board)

					// Create new board from scratch
					my_board = make([][]string, 5)
					board_entry = 0
				} else {
					// Capture line of numbers & add to the board
					board_values := strings.Split(text, " ")
					//board_values := make([]int, len(values))
					//for i, s := range values {
					//	board_values[i], _ = strconv.Atoi(s)
					//}
					my_board[board_entry] = board_values
					board_entry++
				}
			}
			line_number++
		}
	}
	return bingo_numbers, boards
}

func update_board(board [][]string, n string) [][]string {
	//Update a board, replacing its entries by X.
	for i, row := range board {
		for j, value := range row {
			board[i][j] = strings.Replace(value, n, "X", -1)
		}
	}
	fmt.Println(board, "################")
	return board
}

func check_win(board [][]string) bool {
	// Check if a board wins (has 'X' throughout an entire column or row).
	fmt.Println(board[0])
	//win_result := []string{"X", "X", "X", "X", "X"}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != "X" {
				return false
			}
		}
	}
	//if board[0][0] == "X" {
	//fmt.Println("Board has a full column:", board)
	//return true
	//} else if [] {
	//	fmt.Println("Board has a full row:", board)
	//	return true
	//} else {
	//	return false
	//}
	return true
}

func main() {
	// Read data from file
	bingo_numbers, boards := fill_initial_data()

	// What do the boards look like?
	fmt.Println("Boards look like this:", boards)

	// Let's iterate over each number, replacing matching values by '-' in each board
	for _, bingo_num := range bingo_numbers {
		fmt.Println("Bingo number:", bingo_num)
		for _, board := range boards {
			update_board(board, bingo_num)
			win_result := check_win(board)
			if win_result == true {
				fmt.Println("Board has a full column:", board)
				os.Exit(1)
			}
		}
	}
}
