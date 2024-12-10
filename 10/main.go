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

func readInput() [][]int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	input := strings.Split(dataString, "\n")
	input = removeEmptyStrings(input)
	x := make([][]int, 0)
	for _, line := range input {
		lineInt := make([]int, len(line))
		for i, c := range line {
			n, _ := strconv.Atoi(string(c))
			lineInt[i] = n
		}
		x = append(x, lineInt)
	}
	return x
}

func findNeighbor(m [][]int, x int, y int, x0 int, y0 int) map[[4]int]int {
	neighbor := make(map[[4]int]int)
	a, b := x, y-1
	if b >= 0 {
		neighbor[[4]int{x0, y0, a, b}] = m[a][b]
	}
	a, b = x-1, y
	if a >= 0 {
		neighbor[[4]int{x0, y0, a, b}] = m[a][b]
	}
	a, b = x+1, y
	if a < len(m) {
		neighbor[[4]int{x0, y0, a, b}] = m[a][b]
	}
	a, b = x, y+1
	if b < len(m[0]) {
		neighbor[[4]int{x0, y0, a, b}] = m[a][b]
	}
	return neighbor
}

func findStart(m [][]int) map[[4]int]int {
	startPoint := make(map[[4]int]int)
	for x, line := range m {
		for y, n := range line {
			if n == 0 {
				startPoint[[4]int{x, y, x, y}] = 0
			}
		}
	}
	return startPoint
}

func findNext(startPoint map[[4]int]int, m [][]int) map[[4]int]int {
	nextPoints := make(map[[4]int]int)
	var x0, y0, x, y int
	for key, currentPoint := range startPoint {
		x0, y0, x, y = key[0], key[1], key[2], key[3]
		neighbor := findNeighbor(m, x, y, x0, y0)
		for neighborKey, neighborValue := range neighbor {
			if neighborValue == currentPoint+1 {
				nextPoints[neighborKey] = currentPoint + 1
			}
		}
	}
	return nextPoints
}

func findNewNeighbor(m [][]int, x int, y int) map[[2]int]int {
	neighbor := make(map[[2]int]int)
	a, b := x, y-1
	if b >= 0 {
		neighbor[[2]int{a, b}] = m[a][b]
	}
	a, b = x-1, y
	if a >= 0 {
		neighbor[[2]int{a, b}] = m[a][b]
	}
	a, b = x+1, y
	if a < len(m) {
		neighbor[[2]int{a, b}] = m[a][b]
	}
	a, b = x, y+1
	if b < len(m[0]) {
		neighbor[[2]int{a, b}] = m[a][b]
	}
	return neighbor
}

func findNewStart(m [][]int) map[[2]int][2]int {
	startPoint := make(map[[2]int][2]int)
	for x, line := range m {
		for y, n := range line {
			if n == 0 {
				startPoint[[2]int{x, y}] = [2]int{0, 1}
			}
		}
	}
	return startPoint
}

func findNewNext(startPoint map[[2]int][2]int, m [][]int) map[[2]int][2]int {
	nextPoints := make(map[[2]int][2]int)
	var x, y, score, currentPoint int
	for key, value := range startPoint {
		x, y = key[0], key[1]
		score = value[1]
		currentPoint = value[0]
		neighbor := findNewNeighbor(m, x, y)
		for neighborKey, neighborValue := range neighbor {
			if neighborValue == currentPoint+1 {
				_, ok := nextPoints[neighborKey]
				if ok {
					nextPoints[neighborKey] = [2]int{currentPoint + 1, nextPoints[neighborKey][1] + score}
				} else {
					nextPoints[neighborKey] = [2]int{currentPoint + 1, score}
				}
			}
		}
	}
	return nextPoints
}

func calcScore(startPoint map[[2]int][2]int) int {
	sum := 0
	for _, value := range startPoint {
		sum += value[1]
	}
	return sum
}

func main() {
	// Struct is like [start_x, start_y, end_x, end_y] = current_altitude
	m := readInput()
	startPoints := findStart(m)
	for i := 1; i < 10; i++ {
		startPoints = findNext(startPoints, m)
	}
	fmt.Printf("Part 1 : %d\n", len(startPoints))

	startNewPoints := findNewStart(m)
	for i := 1; i < 10; i++ {
		startNewPoints = findNewNext(startNewPoints, m)
	}
	fmt.Printf("Part 2 : %d\n", calcScore(startNewPoints))
}
