package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	input := formatInput(file)
	partOne(input)
	partTwo(input)
}

func partOne(input []int) {
	numberOfInc := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			numberOfInc++
		}
	}
	fmt.Println(numberOfInc)

}

func partTwo(input []int) {
	prev := input[0] + input[1] + input[2]
	numberOfInc := 0
	for i := 3; i < len(input); i++ {
		cur := prev - input[i-3] + input[i]
		if cur > prev {
			numberOfInc++
		}
	}

	fmt.Println(numberOfInc)
}

func formatInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		data = append(data, n)
	}

	return data
}
