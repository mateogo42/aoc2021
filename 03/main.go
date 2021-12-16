package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("sample.txt")
	defer file.Close()

	logs := formatInput(file)

	partOne(logs)
	partTwo(logs)
}

func getMostCommonBitOfColumn(logs []string, col int) byte {
	counter := 0
	for j := 0; j < len(logs); j++ {
		if logs[j][col] == '1' {
			counter++
		} else {
			counter--
		}
	}
	if counter >= 0 {
		return '1'
	}
	return '0'
}

func getLessCommonBitOfColumn(logs []string, col int) byte {
	if getMostCommonBitOfColumn(logs, col) == '1' {
		return '0'
	}

	return '1'
}

func partOne(logs []string) {
	var gamma_arr, epsilon_arr []string
	for i := 0; i < len(logs[0]); i++ {
		mostCommon := getMostCommonBitOfColumn(logs, i)
		lessCommon := getLessCommonBitOfColumn(logs, i)
		gamma_arr = append(gamma_arr, string(mostCommon))
		epsilon_arr = append(epsilon_arr, string(lessCommon))
	}
	gamma, _ := strconv.ParseInt(strings.Join(gamma_arr, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilon_arr, ""), 2, 64)
	fmt.Println(gamma * epsilon)
}

func filterLogs(logs []string, col int, value byte) []string {
	var result []string
	for _, log := range logs {
		if log[col] == value {
			result = append(result, log)
		}
	}

	return result
}

func partTwo(logs []string) {
	var o2, co2 int64
	curr := logs
	col := 0
	for {
		mostCommon := getMostCommonBitOfColumn(curr, col)
		curr = filterLogs(curr, col, mostCommon)
		if len(curr) == 1 {
			o2, _ = strconv.ParseInt(curr[0], 2, 64)
			break
		}
		col++
	}

	curr = logs
	col = 0
	for {
		lessCommon := getLessCommonBitOfColumn(curr, col)
		curr = filterLogs(curr, col, lessCommon)
		if len(curr) == 1 {
			co2, _ = strconv.ParseInt(curr[0], 2, 64)
			break
		}
		col++
	}

	fmt.Println(o2 * co2)
}

func formatInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var logs []string

	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	return logs
}
