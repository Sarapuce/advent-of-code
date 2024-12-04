package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() []string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	grid := strings.Split(dataString, "\n")
	return grid[:len(grid)-1]
}

func findXMas(grid []string) int {
	sum := 0
	for i, line := range grid {
		for j := 0; j < len(line); j++ {
			if j+3 < len(line) {
				horizontal := string(grid[i][j]) + string(grid[i][j+1]) + string(grid[i][j+2]) + string(grid[i][j+3])
				if horizontal == "XMAS" || horizontal == "SAMX" {
					sum++
				}
			}
			if i+3 < len(grid) {
				vertical := string(grid[i][j]) + string(grid[i+1][j]) + string(grid[i+2][j]) + string(grid[i+3][j])
				if vertical == "XMAS" || vertical == "SAMX" {
					sum++
				}
			}
			if i+3 < len(grid) && j+3 < len(line) {
				diagonal := string(grid[i][j]) + string(grid[i+1][j+1]) + string(grid[i+2][j+2]) + string(grid[i+3][j+3])
				if diagonal == "XMAS" || diagonal == "SAMX" {
					sum++
				}
			}
			if i+3 < len(grid) && j-3 >= 0 {
				diagonal := string(grid[i][j]) + string(grid[i+1][j-1]) + string(grid[i+2][j-2]) + string(grid[i+3][j-3])
				if diagonal == "XMAS" || diagonal == "SAMX" {
					sum++
				}
			}
		}
	}
	return sum
}

func findMas(grid []string) int {
	sum := 0
	for i, line := range grid {
		for j := 0; j < len(line); j++ {
			if j+2 < len(line) && i+2 < len(grid) {
				diagonal1 := string(grid[i][j]) + string(grid[i+1][j+1]) + string(grid[i+2][j+2])
				diagonal2 := string(grid[i+2][j]) + string(grid[i+1][j+1]) + string(grid[i][j+2])
				if diagonal1 == "MAS" || diagonal1 == "SAM" {
					if diagonal2 == "MAS" || diagonal2 == "SAM" {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func main() {
	grid := readInput()
	xmasNumber := findXMas(grid)
	fmt.Printf("Part 1 : %d\n", xmasNumber)

	masNumber := findMas(grid)
	fmt.Printf("Part 2 : %d\n", masNumber)
}
