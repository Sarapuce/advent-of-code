package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var memory = make(map[[2]int]int)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func readInput() []int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	dataString = strings.Trim(dataString, "\n")
	input := strings.Split(dataString, " ")
	input = removeEmptyStrings(input)
	stones := make([]int, len(input))
	for i, stoneString := range input {
		n, _ := strconv.Atoi(stoneString)
		stones[i] = n
	}
	return stones
}

func countDigit(n int) int {
	if n == 0 {
		return 0
	}
	c := 0
	for n > 0 {
		n = n / 10
		c++
	}
	return c
}

func power10(exp int) int {
	powersOf10 := []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000}
	if exp >= 0 && exp < len(powersOf10) {
		return powersOf10[exp]
	}
	return -1
}

func splitStone(stone int) (bool, int, int) {
	n := countDigit(stone)
	if n%2 == 0 {
		power := power10(n / 2)
		if power == -1 {
			fmt.Println("Error, number to big")
		}
		a := stone / power
		b := stone % power
		return true, a, b
	} else {
		return false, 0, 0
	}
}

func blink(ogStones []int) []int {
	newStones := make([]int, 0)
	for _, stone := range ogStones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if splitNeeded, a, b := splitStone(stone); splitNeeded {
			newStones = append(newStones, a, b)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}

func newBlink(n int, stone int) int {
	if n == 0 {
		return 1
	} else if s, ok := memory[[2]int{n, stone}]; ok {
		return s
	} else if stone == 0 {
		nextBlink := newBlink(n-1, 1)
		memory[[2]int{n, stone}] = nextBlink
		return nextBlink
	} else if splitNeeded, a, b := splitStone(stone); splitNeeded {
		nextBlink := newBlink(n-1, a) + newBlink(n-1, b)
		memory[[2]int{n, stone}] = nextBlink
		return nextBlink
	} else {
		nextBlink := newBlink(n-1, stone*2024)
		memory[[2]int{n, stone}] = nextBlink
		return nextBlink
	}
}

func main() {
	stones := readInput()
	blinked := make([]int, len(stones))
	copy(blinked, stones)
	for i := 0; i < 25; i++ {
		blinked = blink(blinked)
	}

	fmt.Printf("Part 1 : %d\n", len(blinked))

	sum := 0
	for _, stone := range stones {
		for i := 0; i < 75; i++ {
			newBlink(i, stone)
		}
		sum += newBlink(75, stone)
	}
	fmt.Printf("Part 2 : %d\n", sum)
}
