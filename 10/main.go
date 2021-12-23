package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var pairs map[byte]byte = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	file, _ := os.Open("input.txt")
	lines := formatInput(file)
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	points := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var illegals []byte
	for _, line := range lines {
		stack := make([]byte, len(line)/2)
		for i := 0; i < len(line); i++ {
			if line[i] == '(' || line[i] == '[' || line[i] == '{' || line[i] == '<' {
				stack = append(stack, line[i])
			} else {
				opening := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if !matches(opening, line[i]) {
					illegals = append(illegals, line[i])
				}
			}
		}
	}
	score := 0
	for _, c := range illegals {
		score += points[c]
	}
	fmt.Println(score)
}

func partTwo(lines []string) {
	points := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var scores []int
	for _, line := range lines {
		var stack []byte
		isValid := true
		for i := 0; i < len(line); i++ {
			if line[i] == '(' || line[i] == '[' || line[i] == '{' || line[i] == '<' {
				stack = append(stack, line[i])
			} else {
				opening := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if !matches(opening, line[i]) {
					isValid = false
					break
				}
			}
		}
		if isValid {
			score := 0
			for i := 0; i < len(stack); i++ {
				score = score*5 + points[pairs[stack[len(stack)-1-i]]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
func matches(opening, closing byte) bool {
	if opening == '(' && closing == ')' {
		return true
	}
	if opening == '[' && closing == ']' {
		return true
	}
	if opening == '{' && closing == '}' {
		return true
	}
	if opening == '<' && closing == '>' {
		return true
	}
	return false
}

func formatInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
