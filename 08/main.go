package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	input, output := formatInput(file)
	partOne(output)
	partTwo(input, output)

}

func findByLength(strings []string, length int, numbersMap map[string]int) []string {
	var results []string
	for _, s := range strings {
		if _, found := numbersMap[s]; len(s) == length && !found {
			results = append(results, s)
		}
	}

	return results
}

func findBySegment(line []string, segment string, numbersMap map[string]int) []string {
	var results []string
	for _, l := range line {
		valid := true
		for _, r := range segment {
			if !strings.ContainsRune(string(l), r) {
				valid = false
			}
		}
		if _, found := numbersMap[l]; valid && !found {
			results = append(results, l)
		}
	}

	return results
}

func findByContainer(container string, segments []string, numbersMap map[string]int) []string {
	var results []string
	for _, segment := range segments {
		valid := true
		for _, r := range segment {
			if !strings.ContainsRune(container, r) {
				valid = false
			}
		}
		if _, found := numbersMap[segment]; valid && !found {
			results = append(results, segment)
		}
	}
	return results
}

func partTwo(input [][]string, output [][]string) {
	total := 0
	for i, line := range input {
		var numbersMap = make(map[string]int)
		one := findByLength(line, 2, numbersMap)[0]
		four := findByLength(line, 4, numbersMap)[0]
		seven := findByLength(line, 3, numbersMap)[0]
		eight := findByLength(line, 7, numbersMap)[0]
		numbersMap[one] = 1
		numbersMap[four] = 4
		numbersMap[seven] = 7
		numbersMap[eight] = 8

		segmentsWithLength5 := findByLength(line, 5, numbersMap)
		segmentsWithLength6 := findByLength(line, 6, numbersMap)

		three := findBySegment(segmentsWithLength5, one, numbersMap)[0]
		numbersMap[three] = 3
		nine := findBySegment(segmentsWithLength6, four, numbersMap)[0]
		numbersMap[nine] = 9
		five := findByContainer(nine, segmentsWithLength5, numbersMap)[0]
		numbersMap[five] = 5
		two := findByLength(line, 5, numbersMap)[0]
		numbersMap[two] = 2
		six := findBySegment(segmentsWithLength6, five, numbersMap)[0]
		numbersMap[six] = 6
		zero := findByLength(segmentsWithLength6, 6, numbersMap)[0]
		numbersMap[zero] = 0

		code := 0
		for j := 0; j < len(output[i]); j++ {
			for key, value := range numbersMap {
				if IsEqual(key, output[i][j]) {
					code += value * int(math.Pow10(len(output[i])-j-1))
				}
			}
		}
		total += code
	}
	fmt.Println(total)

}

func IsEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	for _, r := range r1 {
		if !strings.ContainsRune(s2, r) {
			return false
		}
	}
	for _, r := range r2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}

	return true

}

func partOne(output [][]string) {
	count := 0
	for _, line := range output {
		for _, o := range line {
			if len(o) == 2 || len(o) == 3 || len(o) == 4 || len(o) == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func formatInput(file *os.File) ([][]string, [][]string) {
	var input [][]string
	var output [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "|")
		input = append(input, strings.Split(data[0], " "))
		output = append(output, strings.Split(data[1], " "))
	}

	return input, output
}
