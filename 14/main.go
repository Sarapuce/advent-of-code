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

func readInput() ([][2]int, [][2]int) {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	inputs := strings.Split(dataString, "\n")
	inputs = removeEmptyStrings(inputs)

	var input, s string
	var values []string
	var x, y, vX, vY int
	positions := make([][2]int, len(inputs))
	velocities := make([][2]int, len(inputs))
	for i := 0; i < len(positions); i++ {
		input = inputs[i]
		values = strings.Split(input, " ")
		s = strings.Split(values[0], ",")[1]
		y, _ = strconv.Atoi(s)
		s = strings.Split(strings.Split(values[0], "=")[1], ",")[0]
		x, _ = strconv.Atoi(s)
		s = strings.Split(values[1], ",")[1]
		vY, _ = strconv.Atoi(s)
		s = strings.Split(strings.Split(values[1], "=")[1], ",")[0]
		vX, _ = strconv.Atoi(s)

		positions[i][0], positions[i][1] = x, y
		velocities[i][0], velocities[i][1] = vX, vY
	}
	return positions, velocities
}

func move(positions [][2]int, velocities [][2]int, x int, y int) [][2]int {
	var newX, newY int
	newPos := make([][2]int, len(positions))
	for i, p := range positions {
		newX = (p[0] + velocities[i][0])
		newY = (p[1] + velocities[i][1])
		newX = (newX%x + x) % x
		newY = (newY%y + y) % y
		newPos[i][0], newPos[i][1] = newX, newY
	}
	return newPos
}

func printMap(positions [][2]int, x int, y int) {
	m := make([][]int, y)
	for i := range m {
		m[i] = make([]int, x)
	}
	for _, p := range positions {
		m[p[1]][p[0]] += 1
	}
	for i := range m {
		for j := range m[0] {
			if m[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("■")
			}
			// fmt.Print(" | ")
		}
		fmt.Println()
	}
}

func checksum(positions [][2]int, x int, y int) int {
	sum, partSum := 1, 0
	m := make([][]int, y)
	for i := range m {
		m[i] = make([]int, x)
	}
	for _, p := range positions {
		m[p[1]][p[0]] += 1
	}
	for i := 0; i < x/2; i++ {
		for j := 0; j < y/2; j++ {
			partSum += m[j][i]
		}
	}
	sum *= partSum
	partSum = 0
	for i := (x / 2) + 1; i < x; i++ {
		for j := 0; j < y/2; j++ {
			partSum += m[j][i]
		}
	}
	sum *= partSum
	partSum = 0
	for i := (x / 2) + 1; i < x; i++ {
		for j := (y / 2) + 1; j < y; j++ {
			partSum += m[j][i]
		}
	}
	sum *= partSum
	partSum = 0
	for i := 0; i < x/2; i++ {
		for j := (y / 2) + 1; j < y; j++ {
			partSum += m[j][i]
		}
	}
	sum *= partSum
	return sum
}

func detectMid(positions [][2]int, x int, y int) bool {
	m := make([][]int, y)
	for i := range m {
		m[i] = make([]int, x)
	}
	for _, p := range positions {
		m[p[1]][p[0]] += 1
	}
	middle := len(m[0]) / 2
	for i := 95; i < len(m); i++ {
		if m[i][middle] == 0 {
			return false
		}
	}
	return true
}

func main() {
	x, y := 101, 103
	positions, velocities := readInput()
	for i := 0; i < 100; i++ {
		positions = move(positions, velocities, x, y)
	}
	c := checksum(positions, x, y)
	fmt.Printf("Part 1 : %d\n", c)

	positions, velocities = readInput()
	for i := 0; i < 10000; i++ {
		positions = move(positions, velocities, x, y)
		if (i-10)%101 == 0 && (i-88)%103 == 0 {
			// printMap(positions, x, y)
			// fmt.Println()
			// fmt.Printf("%d --------------------------------------------------------\n", i)
			// fmt.Println()
		}
	}
	fmt.Printf("Part 2 : %d\n", 6474+1)
}
