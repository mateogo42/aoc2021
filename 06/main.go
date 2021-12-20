package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fishes map[int]int

func newFishes() Fishes {
	return map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
}

func main() {
	file, _ := os.Open("input.txt")
	fishes := formatInput(file)
	partOne(fishes, 80)
	partOne(fishes, 256)
}

func partOne(fishes map[int]int, days int) {
	for i := 0; i < days; i++ {
		currentFishes := newFishes()
		for counter := range fishes {
			if fishes[counter] > 0 {
				if counter == 0 {
					currentFishes[8] += fishes[0]
					currentFishes[6] += fishes[0]

				} else {
					currentFishes[counter-1] += fishes[counter]
				}
			}
		}
		fishes = currentFishes
	}

	total := 0
	for _, value := range fishes {
		total += value
	}

	fmt.Println(total)
}

func formatInput(file *os.File) map[int]int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fishes := newFishes()

	for _, c := range strings.Split(scanner.Text(), ",") {
		counter, _ := strconv.Atoi(c)
		fishes[counter]++
	}

	return fishes
}
