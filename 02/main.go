package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	linesNumber := strings.Count(dataString, "\n")
	reports := make([][]int, linesNumber)
	for i, line := range strings.Split(dataString, "\n") {
		if line == "" {
			continue
		}
		for _, nStr := range strings.Fields(line) {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				panic("Invalid number in input file")
			}
			reports[i] = append(reports[i], n)
		}
	}
	return reports
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func is_ascending(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] >= numbers[i] {
			return false
		}
	}
	return true
}

func is_descending(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] <= numbers[i] {
			return false
		}
	}
	return true
}

func is_safe(numbers []int) bool {
	ascending := is_ascending(numbers)
	descending := is_descending(numbers)
	for i := 1; i < len(numbers); i++ {
		if abs(numbers[i]-numbers[i-1]) > 3 {
			return false
		}
	}
	return ascending || descending
}

func main() {
	safe_reports := 0
	reports := readInput()
	for _, report := range reports {
		if is_safe(report) {
			safe_reports++
		}
	}
	fmt.Printf("Part 1 : %d\n", safe_reports)

	safe_reports = 0
	for _, report := range reports {
		if is_safe(report) {
			safe_reports++
		} else {
			for j := 0; j < len(report); j++ {
				// Must create a new array at start to avoid modifying report
				fix_report := append(append([]int{}, report[:j]...), report[j+1:]...)
				if is_safe(fix_report) {
					safe_reports++
					break
				}
			}
		}
	}
	fmt.Printf("Part 2 : %d\n", safe_reports)
}
