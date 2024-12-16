package main

import (
	"fmt"
	"os"
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

func printMap(m [][]string) {
	for _, line := range m {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func readInput() [][]string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	inputs := strings.Split(dataString, "\n")
	inputs = removeEmptyStrings(inputs)

	m := make([][]string, len(inputs))
	for i, line := range inputs {
		m[i] = make([]string, len(line))
		for j, c := range line {
			m[i][j] = string(c)
		}
	}

	return m
}

func getNewCoord(x int, y int, d string) (int, int) {
	switch d {
	case "^":
		return x - 1, y
	case "v":
		return x + 1, y
	case ">":
		return x, y + 1
	case "<":
		return x, y - 1
	}
	return -1, -1
}

func getOtherMoves(d string) [3]string {
	switch d {
	case "^":
		return [3]string{">", "<", "v"}
	case ">":
		return [3]string{"^", "<", "v"}
	case "v":
		return [3]string{"^", ">", "<"}
	case "<":
		return [3]string{"^", ">", "v"}
	}
	fmt.Println("getOtherMoves failed")
	return [3]string{"", "", ""}
}

func fillOtherMoves(lab map[[2]int]map[string]int, x int, y int, score int, direction string) {
	otherMoves := getOtherMoves(direction)
	for _, otherMove := range otherMoves {
		if oldScore, ok := lab[[2]int{x, y}][otherMove]; ok {
			if score < oldScore {
				lab[[2]int{x, y}][otherMove] = score
			}
		} else {
			lab[[2]int{x, y}][otherMove] = score
		}
	}
}

func calcNextMove(m [][]string, lab map[[2]int]map[string]int, x int, y int, direction string) {
	var newX, newY, newScore, oldScore int
	currentScore := lab[[2]int{x, y}][direction]
	possibleMoves := make([]string, 0)
	if m[x+1][y] != "#" {
		possibleMoves = append(possibleMoves, "v")
	}
	if m[x-1][y] != "#" {
		possibleMoves = append(possibleMoves, "^")
	}
	if m[x][y+1] != "#" {
		possibleMoves = append(possibleMoves, ">")
	}
	if m[x][y-1] != "#" {
		possibleMoves = append(possibleMoves, "<")
	}

	for _, move := range possibleMoves {
		if move == direction {
			newScore = currentScore + 1
		} else {
			newScore = currentScore + 1001
		}
		newX, newY = getNewCoord(x, y, move)
		if _, ok := lab[[2]int{newX, newY}]; ok {
			if oldScore, ok = lab[[2]int{newX, newY}][move]; ok {
				if newScore < oldScore {
					lab[[2]int{newX, newY}][move] = newScore
					fillOtherMoves(lab, newX, newY, newScore+1000, move)
					calcNextMove(m, lab, newX, newY, move)
				}
			} else {
				lab[[2]int{newX, newY}][direction] = newScore
				calcNextMove(m, lab, newX, newY, direction)
			}
		} else {
			lab[[2]int{newX, newY}] = make(map[string]int)
			lab[[2]int{newX, newY}][move] = newScore
			fillOtherMoves(lab, newX, newY, newScore+1000, move)
			calcNextMove(m, lab, newX, newY, move)
		}
	}
}

func initSolveLab(m [][]string) map[[2]int]map[string]int {
	var x, y int
	lab := make(map[[2]int]map[string]int)
	for i, line := range m {
		for j, c := range line {
			if c == "S" {
				x = i
				y = j
				break
			}
		}
	}
	lab[[2]int{x, y}] = make(map[string]int)
	lab[[2]int{x, y}][">"] = 0
	calcNextMove(m, lab, x, y, ">")
	return lab
}

func getMinimalScore(m [][]string, lab map[[2]int]map[string]int) int {
	var x, y int
	for i, line := range m {
		for j, c := range line {
			if c == "E" {
				x = i
				y = j
			}
		}
	}
	if position, ok := lab[[2]int{x, y}]; ok {
		min := position["^"]
		for _, value := range position {
			if value < min {
				min = value
			}
		}
		return min
	} else {
		fmt.Println("End didn't got calculated")
	}
	return -1
}

func getPreviousCoord(x int, y int, d string) (int, int) {
	switch d {
	case "^":
		return x + 1, y
	case ">":
		return x, y - 1
	case "<":
		return x, y + 1
	case "v":
		return x - 1, y
	}
	fmt.Println("Can't execute getPreviousCoord")
	return -1, -1
}

func findCorrectScore(x int, y int, nextX int, nextY int) string {
	if x == nextX+1 {
		return "^"
	} else if x == nextX-1 {
		return "v"
	} else if y == nextY+1 {
		return "<"
	} else if y == nextY-1 {
		return ">"
	}
	return ""
}

func getInitialDirection(m [][]string, lab map[[2]int]map[string]int) (int, int, [][2]int) {
	var x, y int
	for i, line := range m {
		for j, c := range line {
			if c == "E" {
				x = i
				y = j
				break
			}
		}
	}

	previousCoord := make([][2]int, 0)
	miniScore := lab[[2]int{x, y}]["^"]
	for _, score := range lab[[2]int{x, y}] {
		if score < miniScore {
			miniScore = score
		}
	}

	for key, score := range lab[[2]int{x, y}] {
		if score == miniScore {
			switch key {
			case "^":
				previousCoord = append(previousCoord, [2]int{x + 1, y})
			case "v":
				previousCoord = append(previousCoord, [2]int{x - 1, y})
			case ">":
				previousCoord = append(previousCoord, [2]int{x, y - 1})
			case "<":
				previousCoord = append(previousCoord, [2]int{x, y + 1})
			}
		}
	}
	return x, y, previousCoord
}

func getMinimalPathTo(m [][]string, lab map[[2]int]map[string]int, x int, y int, nextX int, nextY int) {
	var oldX, oldY, gapScore int
	minimalDirection := findCorrectScore(x, y, nextX, nextY)
	currentScore := lab[[2]int{x, y}][minimalDirection]
	m[x][y] = "O"

	possibleMoves := make([]string, 0)
	if m[x+1][y] != "#" {
		possibleMoves = append(possibleMoves, "^")
	}
	if m[x-1][y] != "#" {
		possibleMoves = append(possibleMoves, "v")
	}
	if m[x][y+1] != "#" {
		possibleMoves = append(possibleMoves, "<")
	}
	if m[x][y-1] != "#" {
		possibleMoves = append(possibleMoves, ">")
	}

	for _, possibleMove := range possibleMoves {
		if possibleMove == minimalDirection {
			gapScore = 1
		} else {
			gapScore = 1001
		}
		oldX, oldY = getPreviousCoord(x, y, possibleMove)
		if previousScore, ok := lab[[2]int{oldX, oldY}][possibleMove]; ok {
			if previousScore == currentScore-gapScore {
				m[oldX][oldY] = "O"
				getMinimalPathTo(m, lab, oldX, oldY, x, y)
			}
		} else {
			fmt.Println("Previous step not calculated")
		}
	}

}

func checksum(m [][]string) int {
	sum := 0
	for _, line := range m {
		for _, c := range line {
			if c == "O" || c == "E" || c == "S" {
				sum++
			}
		}
	}
	return sum
}

func main() {
	m := readInput()
	lab := initSolveLab(m)
	fmt.Printf("Part 1 : %d\n", getMinimalScore(m, lab))

	x, y, previousCoords := getInitialDirection(m, lab)
	for _, previousCoord := range previousCoords {
		getMinimalPathTo(m, lab, previousCoord[0], previousCoord[1], x, y)
	}
	fmt.Printf("Part 2 : %d\n", checksum(m))
}
