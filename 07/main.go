package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func readInput() ([]int, [][]int) {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	input := strings.Split(dataString, "\n")
	input = removeEmptyStrings(input)
	results := make([]int, len(input))
	numbers := make([][]int, len(input))
	for i, line := range input {
		line := strings.Split(line, ":")
		result, _ := strconv.Atoi(line[0])
		results[i] = result
		number := removeEmptyStrings(strings.Split(line[1], " "))
		numbers[i] = make([]int, len(number))
		for j, n := range number {
			if n != "" {
				num, _ := strconv.Atoi(n)
				numbers[i][j] = num
			}
		}
	}
	return results, numbers
}

func getPossibilities(result int, numbers []int) []int {
	possibilities := []int{numbers[0]}
	for j := 1; j < len(numbers); j++ {
		newPossibilities := []int{}
		for k := 0; k < len(possibilities); k++ {
			multiplicationResult := possibilities[k] * numbers[j]
			additionResult := possibilities[k] + numbers[j]
			if multiplicationResult <= result {
				newPossibilities = append(newPossibilities, multiplicationResult)
			}
			if additionResult <= result {
				newPossibilities = append(newPossibilities, additionResult)
			}
		}
		possibilities = newPossibilities
	}
	return possibilities
}

func getNewPossibilities(result int, numbers []int) []int {
	possibilities := []int{numbers[0]}
	for j := 1; j < len(numbers); j++ {
		newPossibilities := []int{}
		for k := 0; k < len(possibilities); k++ {
			multiplicationResult := possibilities[k] * numbers[j]
			additionResult := possibilities[k] + numbers[j]
			concatenationResult := concatenate(possibilities[k], numbers[j])
			if multiplicationResult <= result {
				newPossibilities = append(newPossibilities, multiplicationResult)
			}
			if additionResult <= result {
				newPossibilities = append(newPossibilities, additionResult)
			}
			if concatenationResult <= result {
				newPossibilities = append(newPossibilities, concatenationResult)
			}
		}
		possibilities = newPossibilities
	}
	return possibilities
}

func resInPossibilities(result int, possibilities []int) bool {
	for _, p := range possibilities {
		if p == result {
			return true
		}
	}
	return false
}

func concatenate(a int, b int) int { // Probably slower than using number operation but it works
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)
	n, _ := strconv.Atoi(sa + sb)
	return n
}

func main() {
	results, numbers := readInput()
	sum := 0
	for i, result := range results {
		if resInPossibilities(result, getPossibilities(result, numbers[i])) {
			sum += result
		}
	}
	fmt.Printf("Part 1 : %d\n", sum)

	sum = 0
	for i, result := range results {
		if resInPossibilities(result, getNewPossibilities(result, numbers[i])) {
			sum += result
		}
	}
	fmt.Printf("Part 2 : %d\n", sum)
}
