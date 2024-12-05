package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([][]int, [][]int) {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	input := strings.Split(dataString, "\n\n")
	rules_string := strings.Split(input[0], "\n")
	manuals_string := strings.Split(input[1], "\n")
	rules_string, manuals_string = removeEmptyStrings(rules_string), removeEmptyStrings(manuals_string)

	rules := make([][]int, len(rules_string))
	for i, rule := range rules_string {
		rules[i] = make([]int, 2)
		rule_split := strings.Split(rule, "|")
		n1, _ := strconv.Atoi(rule_split[0])
		n2, _ := strconv.Atoi(rule_split[1])
		rules[i][0], rules[i][1] = n1, n2
	}

	manuals := make([][]int, len(manuals_string))
	for i, manual := range manuals_string {
		pages := strings.Split(manual, ",")
		manuals[i] = make([]int, len(pages))
		for j, nStr := range pages {
			n, _ := strconv.Atoi(nStr)
			manuals[i][j] = n
		}
	}
	return rules, manuals
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func checkManual(manual []int, rules [][]int) int {
	for i := 0; i < len(manual)-1; i++ {
		for _, rule := range rules {
			if rule[0] == manual[i] && rule[1] == manual[i+1] {
				break
			}
			if rule[0] == manual[i+1] && rule[1] == manual[i] {
				return 0
			}
		}
	}
	return manual[(len(manual)-1)/2]
}

func selectIncorrectManuals(manuals [][]int, rules [][]int) [][]int {
	incorrectManuals := make([][]int, 0)
	for _, manual := range manuals {
		if checkManual(manual, rules) == 0 { // No zero in the inputs so it can be used as condition to select incorrect manuals
			incorrectManuals = append(incorrectManuals, manual)
		}
	}
	return incorrectManuals
}

func reorderManuals(manual []int, rules [][]int) ([]int, bool) {
	for i := 0; i < len(manual)-1; i++ {
		for _, rule := range rules {
			if rule[0] == manual[i+1] && rule[1] == manual[i] {
				manual[i], manual[i+1] = manual[i+1], manual[i]
				return manual, false
			}
		}
	}
	return manual, true
}

func main() {
	rules, manuals := readInput()
	sum := 0
	for _, manual := range manuals {
		sum += checkManual(manual, rules)
	}
	fmt.Printf("Part 1 : %d\n", sum)

	sum = 0
	incorrectManuals := selectIncorrectManuals(manuals, rules)
	for _, manual := range incorrectManuals {
		sorted := false
		for !sorted {
			manual, sorted = reorderManuals(manual, rules)
		}
		sum += manual[(len(manual)-1)/2]
	}
	fmt.Printf("Part 2 : %d\n", sum)
}
