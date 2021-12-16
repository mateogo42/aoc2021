package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	numbers, boards := formatInput(file)
	partOne(numbers, boards)
	partTwo(numbers, boards)
}

func partOne(numbers []int, boards [][]int) {
	markedBoards := make([][25]bool, len(boards))
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(boards); j++ {
			// Mark the numbers in each board
			markedBoards[j] = markBoard(numbers[i], boards[j], markedBoards[j])
			// Check if the current board has already completed a row or a column.
			if checkBoard(markedBoards[j]) {
				// If it has, we have a winner, calculate the score for that board
				unmarked := 0
				for k := 0; k < len(boards[j]); k++ {
					if !markedBoards[j][k] {
						unmarked = unmarked + boards[j][k]
					}
				}
				fmt.Println(unmarked * numbers[i])
				return

			}
		}

	}
}

func partTwo(numbers []int, boards [][]int) {
	markedBoards := make([][25]bool, len(boards))
	finishedBoards := make(map[int]struct{})
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(boards); j++ {
			// Mark the number in each board
			markedBoards[j] = markBoard(numbers[i], boards[j], markedBoards[j])
			// Check if the current board has already completed a row or a column.
			if _, ok := finishedBoards[j]; !ok && checkBoard(markedBoards[j]) {
				// Added to the map of completed boards
				finishedBoards[j] = struct{}{}
				// If it is the last board to complete a row or a column, we have a winner, calculate the score
				if len(finishedBoards) == len(boards) {
					unmarked := 0
					for k := 0; k < len(boards[j]); k++ {
						if !markedBoards[j][k] {
							unmarked = unmarked + boards[j][k]
						}
					}
					fmt.Println(unmarked * numbers[i])
					return
				}

			}
		}

	}

}

func checkBoard(marked [25]bool) bool {
	// Check columns
	for i := 0; i < 5; i++ {
		winCol := true
		for j := i; j < 25; j = j + 5 {
			winCol = winCol && marked[j]
		}
		if winCol {
			return true
		}
	}

	//Check rows
	for i := 0; i < 25; i = i + 5 {
		winRow := true
		for j := 0; j < 5; j++ {
			winRow = winRow && marked[i+j]
		}
		if winRow {
			return true
		}
	}

	return false
}

func markBoard(number int, board []int, marked [25]bool) [25]bool {
	for i := 0; i < len(board); i++ {
		if board[i] == number {
			marked[i] = true
		}
	}
	return marked
}

func formatInput(file *os.File) ([]int, [][]int) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var numbers []int
	var boards [][]int
	for _, num := range strings.Split(scanner.Text(), ",") {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}
	scanner.Scan()
	var board []int
	var counter int
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		for _, num := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(num)
			board = append(board, n)
		}
		counter++
		if counter == 5 {
			boards = append(boards, board)
			board = []int{}
			counter = 0
		}
	}

	return numbers, boards
}
