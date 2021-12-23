package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LowPoint struct {
	x     int
	y     int
	value int
}

func main() {
	file, _ := os.Open("input.txt")
	heightMap := formatInput(file)
	partOne(heightMap)
	partTwo(heightMap)

}

func partTwo(heightMap [][]int) {
	var basins []int
	visited := make([][]bool, len(heightMap))
	for i := range heightMap {
		visited[i] = make([]bool, len(heightMap[i]))
	}
	lowPoints := findLowPoints(heightMap)
	var basinSize int
	for _, point := range lowPoints {
		basinSize, visited = checkBasin(point.x, point.y, heightMap, visited)
		basins = append(basins, basinSize)
	}
	sort.Ints(basins)
	total := 1
	for i := 0; i < 3; i++ {
		total *= basins[len(basins)-1-i]
	}
	fmt.Println(total)
}

func checkBasin(i, j int, heightMap [][]int, visited [][]bool) (int, [][]bool) {
	if i < 0 || j < 0 || i > len(heightMap)-1 || j > len(heightMap[i])-1 {
		return 0, visited
	}
	if visited[i][j] {
		return 0, visited
	}
	visited[i][j] = true
	if heightMap[i][j] == 9 {
		return 0, visited
	}
	up, visited := checkBasin(i-1, j, heightMap, visited)
	down, visited := checkBasin(i+1, j, heightMap, visited)
	left, visited := checkBasin(i, j-1, heightMap, visited)
	right, visited := checkBasin(i, j+1, heightMap, visited)
	return 1 + up + down + left + right, visited
}

func findLowPoints(heightMap [][]int) []LowPoint {
	var lowPoints []LowPoint

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			isLowPoint := true
			if i > 0 {
				isLowPoint = isLowPoint && heightMap[i][j] < heightMap[i-1][j]
			}
			if j > 0 {
				isLowPoint = isLowPoint && heightMap[i][j] < heightMap[i][j-1]
			}
			if i < len(heightMap)-1 {
				isLowPoint = isLowPoint && heightMap[i][j] < heightMap[i+1][j]
			}
			if j < len(heightMap[i])-1 {
				isLowPoint = isLowPoint && heightMap[i][j] < heightMap[i][j+1]
			}
			if isLowPoint {
				lowPoints = append(lowPoints, LowPoint{x: i, y: j, value: heightMap[i][j]})
			}
		}
	}
	return lowPoints
}

func partOne(heightMap [][]int) {
	lowPoints := findLowPoints(heightMap)
	riskLevel := 0
	for _, low := range lowPoints {
		riskLevel += low.value
	}
	fmt.Println(riskLevel)
}

func formatInput(file *os.File) [][]int {
	var heightMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "")
		var level []int
		for i := 0; i < len(nums); i++ {
			h, _ := strconv.Atoi(nums[i])
			level = append(level, h)
		}
		heightMap = append(heightMap, level)
	}

	return heightMap
}
