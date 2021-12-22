package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	nums := formatInput(file)
	partOne(nums)
	partTwo(nums)
}

func partOne(nums []int) {
	sort.Ints(nums)
	n := len(nums)
	var median float64
	if n%2 == 0 {
		median = float64((nums[n/2] + nums[n/2-1]) / 2)
	} else {
		median = float64(nums[n/2])
	}

	fuel := 0.0

	for i := 0; i < n; i++ {
		fuel += math.Abs(median - float64(nums[i]))
	}
	fmt.Println(fuel)
}

func partTwo(nums []int) {
	sum := 0.0
	n := float64(len(nums))
	for _, num := range nums {
		sum += float64(num)
	}
	mean := int(sum / n)

	fuel := 0
	for i := 0; i < int(n); i++ {
		m := int(math.Abs(float64(mean) - float64(nums[i])))
		fuel += m * (m + 1) / 2
	}
	fmt.Println(fuel)
}

func formatInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var nums []int

	values := strings.Split(scanner.Text(), ",")

	for _, val := range values {
		n, _ := strconv.Atoi(val)
		nums = append(nums, n)
	}

	return nums

}
