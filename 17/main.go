package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() (int, int, int, []int) {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	inputs := strings.Split(dataString, "\n\n")

	registerStrings := strings.Split(inputs[0], "\n")
	aString := strings.Split(registerStrings[0], ": ")[1]
	bString := strings.Split(registerStrings[1], ": ")[1]
	cString := strings.Split(registerStrings[2], ": ")[1]
	a, _ := strconv.Atoi(aString)
	b, _ := strconv.Atoi(bString)
	c, _ := strconv.Atoi(cString)

	programString := strings.Split(inputs[1], ": ")[1]
	programList := strings.Split(programString, ",")
	program := make([]int, len(programList))
	for i, instruction := range programList {
		program[i], _ = strconv.Atoi(instruction)
	}

	return a, b, c, program
}

func power2(exp int) int {
	powersOf2 := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	if exp >= 0 && exp < len(powersOf2) {
		return powersOf2[exp]
	}
	return 0
}

func getComboValue(a int, b int, c int, operand int) int {
	if operand <= 3 {
		return operand
	} else {
		switch operand {
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		}
	}
	return -1
}

func adv(a int, combo int) int {
	return a / power2(combo)
}

func bxl(b int, operand int) int {
	return b ^ operand
}

func bst(combo int) int {
	return combo % 8
}

func jnz(a int) bool {
	return a != 0
}

func bxc(b int, c int) int {
	return b ^ c
}

func out(combo int) int {
	return combo % 8
}

func bdv(a int, combo int) int {
	return a / power2(combo)
}

func cdv(a int, combo int) int {
	return a / power2(combo)
}

func executeProgram(a int, b int, c int, program []int) []int {
	output := make([]int, 0)
	var opcode, operand, combo int
	for rip := 0; rip < len(program); rip = rip + 2 {
		opcode = program[rip]
		operand = program[rip+1]
		combo = getComboValue(a, b, c, operand)
		switch opcode {
		case 0:
			a = adv(a, combo)
		case 1:
			b = bxl(b, operand)
		case 2:
			b = bst(combo)
		case 3:
			if jnz(a) {
				rip = operand
				rip -= 2
			}
		case 4:
			b = bxc(b, c)
		case 5:
			output = append(output, out(combo))
		case 6:
			b = bdv(a, combo)
		case 7:
			c = cdv(a, combo)
		}
	}
	return output
}

func getStringOutput(output []int) string {
	var outputString []string
	for _, val := range output {
		outputString = append(outputString, strconv.Itoa(val))
	}
	return strings.Join(outputString, ",")
}

func getBase8(x int) string {
	s := strconv.FormatInt(int64(x), 8)
	return s
}

func OctalStringToInt(octalStr string) int {
	n, _ := strconv.ParseInt(octalStr, 8, 64)
	return int(n)
}

func findNumber(b int, c int, program []int) []string { // Change 1 digit in the base 8 number change one digit of the output. This bruteforce to get all possibilities
	s := make([]string, 0)
	s = append(s, strings.Repeat("1", len(program)))
	n := len(program)
	var jStr, candidate string
	var a int
	var output []int
	for i := 0; i < n; i++ {
		nextS := make([]string, 0)
		for _, candidate = range s {
			for j := 0; j < 8; j++ {
				jStr = strconv.FormatInt(int64(j), 10)
				candidate = candidate[:i] + jStr + candidate[i+1:]
				a = OctalStringToInt(candidate)
				output = executeProgram(a, b, c, program)
				if len(output) == n && output[n-1-i] == program[n-1-i] {
					nextS = append(nextS, candidate)
				}
			}
		}
		s = nextS
	}
	return s
}

func main() {
	a, b, c, program := readInput()

	output := executeProgram(a, b, c, program)
	fmt.Printf("Part 1 : %s\n", getStringOutput(output))

	candidates := findNumber(b, c, program)
	fmt.Printf("Part 2 : %d\n", OctalStringToInt(candidates[0]))
}
