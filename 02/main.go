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
	defer file.Close()

	commands, values := formatInput(file)
	partOne(commands, values)
	partTwo(commands, values)
}

func partOne(commands []string, values []int) {
	x := 0 // Horizontal position
	y := 0 // Vertical position
	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "forward":
			x += values[i]
		case "down":
			y += values[i]
		case "up":
			y -= values[i]
		}
	}
	fmt.Println(x * y)
}
func partTwo(commands []string, values []int) {
	x := 0 // Horizontal position
	y := 0 // Vertical position
	aim := 0
	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "forward":
			x += values[i]
			y += values[i] * aim
		case "down":
			aim += values[i]
		case "up":
			aim -= values[i]
		}
	}
	fmt.Println(x * y)

}

func formatInput(file *os.File) ([]string, []int) {
	scanner := bufio.NewScanner(file)
	var commands []string
	var values []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(line[1])
		commands = append(commands, line[0])
		values = append(values, value)
	}

	return commands, values

}
