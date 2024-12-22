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

func readInput() [][2]int {
	fileName := "input.txt"
	var elements []string
	var x, y int
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	inputs := strings.Split(dataString, "\n")
	inputs = removeEmptyStrings(inputs)

	wrongBytes := make([][2]int, len(inputs))
	for i, input := range inputs {
		elements = strings.Split(input, ",")
		x, _ = strconv.Atoi(elements[0])
		y, _ = strconv.Atoi(elements[1])
		wrongBytes[i][0], wrongBytes[i][1] = x, y
	}

	return wrongBytes
}

func printMap(m [][]string) {
	for _, line := range m {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func createMaze(wrongBytes [][2]int, size int) [][]string {
	maze := make([][]string, size)
	for i := 0; i < size; i++ {
		maze[i] = make([]string, size)
		for j := 0; j < size; j++ {
			maze[i][j] = "."
		}
	}
	var x, y int
	for _, wrongByte := range wrongBytes {
		x, y = wrongByte[0], wrongByte[1]
		maze[y][x] = "#"
	}
	maze[0][0] = "0"
	return maze
}

func getPossible(x int, y int, maze [][]string) [][2]int {
	possibility := make([][2]int, 0)
	if x-1 >= 0 && maze[y][x-1] != "#" {
		possibility = append(possibility, [2]int{y, x - 1})
	}
	if x+1 < len(maze) && maze[y][x+1] != "#" {
		possibility = append(possibility, [2]int{y, x + 1})
	}
	if y-1 >= 0 && maze[y-1][x] != "#" {
		possibility = append(possibility, [2]int{y - 1, x})
	}
	if y+1 < len(maze) && maze[y+1][x] != "#" {
		possibility = append(possibility, [2]int{y + 1, x})
	}
	return possibility
}

func fillCase(x int, y int, maze [][]string) {
	if maze[y][x] != "." {
		return
	}
	var xx, yy, value int
	minimal := 70 * 70
	possibility := getPossible(x, y, maze)
	nextCase := make([][2]int, 0)
	for _, possibility := range possibility {
		xx, yy = possibility[0], possibility[1]
		value64, err := strconv.ParseInt(maze[yy][xx], 10, 64)
		if err == nil {
			value = int(value64)
			if value < minimal {
				minimal = value
			}
		} else {
			nextCase = append(nextCase, [2]int{xx, yy})
		}
	}
	newValueStr := strconv.FormatInt(int64(minimal+1), 10)
	maze[y][x] = newValueStr
	for _, next := range nextCase {
		xx, yy = next[0], next[1]
		fillCase(xx, yy, maze)
	}
}

func main() {
	wrongBytes := readInput()
	maze := createMaze(wrongBytes, 7)
	fillCase(0, 1, maze)
	fillCase(1, 0, maze)
	fmt.Println(maze)
	fmt.Printf("Part 1 : %d\n", 0)

	fmt.Printf("Part 2 : %d\n", 0)
}
