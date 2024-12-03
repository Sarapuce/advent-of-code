package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	return dataString
}

func getFirstPart(s string) []string {
	return strings.Split(s, ")")
}

func getLastPart(s string) string {
	s_part := strings.Split(s, "mul(")
	return s_part[len(s_part)-1]
}

func validateNumber(input string) bool {
	pattern := `^\d{1,3},\d{1,3}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(input)
}

func mul(s string) int {
	nums := strings.Split(s, ",")
	num1, num2 := nums[0], nums[1]
	num1_int, _ := strconv.Atoi(num1)
	num2_int, _ := strconv.Atoi(num2)
	return num1_int * num2_int
}

func computeSum(s string) int {
	sum := 0
	factor := 1
	l := len(s)
	for i := 0; i < l-8; i++ {
		if s[i:i+4] == "do()" {
			factor = 1
		}
		if s[i:i+7] == "don't()" {
			factor = 0
		}
		for j := i + 8; j < i+13 && j < l; j++ {
			if validateMul(s[i:j]) {
				op := s[i : j-1]
				n := strings.Split(op, "(")[1]
				sum += factor * mul(n)
			}
		}
	}
	return sum
}

func validateMul(input string) bool {
	pattern := `^mul\(\d{1,3},\d{1,3}\)$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(input)
}

func main() {
	input := readInput()
	sum := 0
	stripInput := getFirstPart(input)
	for _, line := range stripInput {
		clean_line := getLastPart(line)
		if validateNumber(clean_line) {
			sum += mul(clean_line)
		}
	}
	fmt.Printf("Part 1 : %d\n", sum)

	sum = computeSum(input)
	fmt.Printf("Part 2 : %d\n", sum)
}
