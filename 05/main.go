package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	starts, ends, maxX, maxY := formatInput(file)
	partOne(starts, ends, maxX, maxY)
	partTwo(starts, ends, maxX, maxY)
}

func partOne(starts [][2]int, ends [][2]int, maxX, maxY int) {
	diagram := make([][]int, maxY)
	for i := range diagram {
		diagram[i] = make([]int, maxX)
	}
	overlaps := 0
	for i := 0; i < len(starts); i++ {
		if starts[i][0] == ends[i][0] || starts[i][1] == ends[i][1] {
			x := starts[i][0]
			y := starts[i][1]
			dirX, dirY := getDirection(starts[i], ends[i])
			for {
				diagram[y][x]++
				if diagram[y][x] == 2 {
					overlaps++
				}
				if x == ends[i][0] && y == ends[i][1] {
					break
				}
				x = x + dirX
				y = y + dirY
			}
		}
	}
	fmt.Println(overlaps)
}

func partTwo(starts [][2]int, ends [][2]int, maxX, maxY int) {
	diagram := make([][]int, maxY)
	for i := range diagram {
		diagram[i] = make([]int, maxX)
	}
	overlaps := 0
	for i := 0; i < len(starts); i++ {
		x := starts[i][0]
		y := starts[i][1]
		dirX, dirY := getDirection(starts[i], ends[i])
		for {
			diagram[y][x]++
			if diagram[y][x] == 2 {
				overlaps++
			}
			if x == ends[i][0] && y == ends[i][1] {
				break
			}
			x = x + dirX
			y = y + dirY
		}
	}
	fmt.Println(overlaps)
}

func getDirection(start [2]int, end [2]int) (int, int) {
	dirX, dirY := 0, 0
	if start[0] < end[0] {
		dirX = 1
	}
	if start[0] > end[0] {
		dirX = -1
	}

	if start[1] < end[1] {
		dirY = 1
	}
	if start[1] > end[1] {
		dirY = -1
	}

	return dirX, dirY
}

func formatInput(file *os.File) ([][2]int, [][2]int, int, int) {
	scanner := bufio.NewScanner(file)
	maxX := 0
	maxY := 0
	var starts [][2]int
	var ends [][2]int
	for scanner.Scan() {
		segment := strings.Split(scanner.Text(), " -> ")
		starts = append(starts, parsePoint(segment[0]))
		ends = append(ends, parsePoint(segment[1]))
	}

	for i := 0; i < len(starts); i++ {
		maxX = int(math.Max(float64(maxX), float64(starts[i][0])))
		maxX = int(math.Max(float64(maxX), float64(ends[i][0])))

		maxY = int(math.Max(float64(maxY), float64(starts[i][1])))
		maxY = int(math.Max(float64(maxY), float64(ends[i][1])))
	}

	return starts, ends, maxX + 1, maxY + 1
}

func parsePoint(point string) [2]int {
	coords := strings.Split(point, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return [2]int{x, y}

}
