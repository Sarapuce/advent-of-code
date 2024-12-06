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
	input := strings.Split(dataString, "\n")
	input = removeEmptyStrings(input)
	return input
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

func countX(world []string) int {
	sum := 0
	for x, line := range world {
		for y := 0; y < len(line); y++ {
			if string(world[x][y]) == "X" {
				sum += 1
			}
		}
	}
	return sum
}

func getPLayerPos(world []string) (int, int, string) {
	for x, line := range world {
		for y := 0; y < len(line); y++ {
			switch string(world[x][y]) {
			case string("^"):
				return x, y, "u"
			case string(">"):
				return x, y, "r"
			case string("<"):
				return x, y, "l"
			case string("v"):
				return x, y, "d"
			}
		}
	}
	return 0, 0, "?"
}

func makeAStep(world []string, x int, y int, direction string) ([]string, int, int, string) {
	switch direction {
	case "u":
		if x == 0 {
			world[x] = world[x][:y] + "X" + world[x][y+1:]
			return world, x - 1, y, "exit"
		} else if string(world[x-1][y]) == "#" {
			world[x] = world[x][:y] + ">" + world[x][y+1:]
			return world, x, y, "r"
		} else {
			world[x] = world[x][:y] + "X" + world[x][y+1:]
			world[x-1] = world[x-1][:y] + "^" + world[x-1][y+1:]
			return world, x - 1, y, "u"
		}
	case "d":
		if x == len(world)-1 {
			world[x] = world[x][:y] + "X" + world[x][y+1:]
			return world, x + 1, y, "exit"
		} else if string(world[x+1][y]) == "#" {
			world[x] = world[x][:y] + "<" + world[x][y+1:]
			return world, x, y, "l"
		} else {
			world[x] = world[x][:y] + "X" + world[x][y+1:]
			world[x+1] = world[x+1][:y] + "v" + world[x+1][y+1:]
			return world, x + 1, y, "d"
		}
	case "r":
		if y == len(world[0])-1 {
			world[x] = world[x][:y] + "X"
			return world, x, y + 1, "exit"
		} else if string(world[x][y+1]) == "#" {
			world[x] = world[x][:y] + "v" + world[x][y+1:]
			return world, x, y, "d"
		} else {
			world[x] = world[x][:y] + "X" + world[x][y+1:]
			world[x] = world[x][:y+1] + ">" + world[x][y+2:]
			return world, x, y + 1, "r"
		}
	case "l":
		if y == 0 {
			world[x] = "X" + world[x][y+1:]
			return world, x, y - 1, "exit"
		} else if string(world[x][y-1]) == "#" {
			world[x] = world[x][:y] + "^" + world[x][y+1:]
			return world, x, y, "u"
		} else {
			world[x] = world[x][:y] + "X" + world[x][y+1:]
			world[x] = world[x][:y-1] + "<" + world[x][y:]
			return world, x, y - 1, "l"
		}
	}
	return world, x, y, "?"
}

// func printWorld(world []string, x int) {
// 	if x-3 < 0 {
// 		x = 3
// 	}
// 	if x+3 > len(world) {
// 		x = len(world) - 3
// 	}
// 	for i := x - 3; i < x+3; i++ {
// 		println(world[i])
// 	}
// }

func main() {
	world := readInput()
	x0, y0, direction := getPLayerPos(world)
	x, y := x0, y0
	for direction != "exit" {
		world, x, y, direction = makeAStep(world, x, y, direction)
		// printWorld(world, x)
	}
	steps := countX(world)
	fmt.Printf("Part 1 : %d\n", steps)

	sum := 0
	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[0]); j++ {
			world = readInput()
			x0, y0, direction := getPLayerPos(world)
			x, y = x0, y0
			if string(world[i][j]) != "#" && !(i == x0 && j == y0) { // Don't put obstacle on already existing object and soldier start
				world[i] = world[i][:j] + "#" + world[i][j+1:]
				stepNumber := 0
				for direction != "exit" && stepNumber < 20000 { // If the number of step is too high, the soldier must be stuck. The correct step number should be higher I think
					world, x, y, direction = makeAStep(world, x, y, direction)
					stepNumber++
				}
				if direction != "exit" {
					sum += 1
				}
			}
		}
	}
	fmt.Printf("Part 2 : %d\n", sum)
}
