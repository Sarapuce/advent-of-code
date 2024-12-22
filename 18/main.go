package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 71

type Point struct {
	kind  string
	value int
}

type coords [2]int

type maze map[coords]Point

func (m *maze) addWall(x int, y int) {
	c := coords{x, y}
	(*m)[c] = Point{"#", -1}
}

func (m *maze) addSurroundingWalls() {
	for i := -1; i < size+1; i++ {
		m.addWall(-1, i)
		m.addWall(size, i)
		m.addWall(i, -1)
		m.addWall(i, size)
	}
}

func (m maze) print() {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			c := coords{x, y}
			point, ok := m[c]
			if !ok {
				fmt.Print(".")
			} else {
				fmt.Print(point.kind)
			}
		}
		fmt.Println()
	}
}

func (m *maze) calculatePoint(c coords) []coords {
	min := 5000 // greater than 70*70
	nextNeighbors := make([]coords, 0)
	neighbors := [4]coords{
		{c[0] + 1, c[1]},
		{c[0], c[1] + 1},
		{c[0] - 1, c[1]},
		{c[0], c[1] - 1}}

	for _, neighbor := range neighbors {
		p, ok := (*m)[neighbor]
		if !ok {
			(*m)[neighbor] = Point{".", -1}
			nextNeighbors = append(nextNeighbors, neighbor)
		} else {
			if p.kind == "." {
				nextNeighbors = append(nextNeighbors, neighbor)
			} else if p.kind == "O" {
				if p.value < min-1 {
					min = p.value + 1
				}
			}
		}
	}
	(*m)[c] = Point{"O", min}
	return nextNeighbors
}

func (m *maze) bootstrapMaze() []coords {
	nextNeighbors := make([]coords, 0)
	(*m)[coords{0, 0}] = Point{"O", 0}
	point, ok := (*m)[coords{1, 0}]
	if !ok {
		(*m)[coords{1, 0}] = Point{".", 0}
		nextNeighbors = append(nextNeighbors, coords{1, 0})
	} else if point.kind != "#" {
		nextNeighbors = append(nextNeighbors, coords{1, 0})
	}

	point, ok = (*m)[coords{0, 1}]
	if !ok {
		(*m)[coords{0, 1}] = Point{".", 0}
		nextNeighbors = append(nextNeighbors, coords{0, 1})
	} else if point.kind != "#" {
		nextNeighbors = append(nextNeighbors, coords{0, 1})
	}
	return nextNeighbors
}

func (m *maze) isSolved() bool {
	end, ok := (*m)[coords{size - 1, size - 1}]
	return ok && end.kind == "O"
}

func (m *maze) getPathLength() int {
	end := (*m)[coords{size - 1, size - 1}]
	return end.value
}

func (m *maze) solve() (int, bool) {
	neighbors := m.bootstrapMaze()
	for solved := m.isSolved(); !solved; solved = m.isSolved() {
		nextNeighbord := make([]coords, 0)
		for _, neighbor := range neighbors {
			nextNeighbord = append(nextNeighbord, m.calculatePoint(neighbor)...)
		}
		nextNeighbord = removeDuplicates(nextNeighbord)
		neighbors = nextNeighbord
		if len(neighbors) == 0 && !m.isSolved() {
			return -1, false
		}
	}
	return m.getPathLength(), true
}

func removeDuplicates(slice []coords) []coords {
	seen := make(map[coords]bool)
	result := []coords{}

	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
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

func readInput() []coords {
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

	walls := make([]coords, len(inputs))
	for i, input := range inputs {
		elements = strings.Split(input, ",")
		x, _ = strconv.Atoi(elements[0])
		y, _ = strconv.Atoi(elements[1])
		walls[i] = coords{x, y}
	}

	return walls
}

func createMaze(inputs []coords, wallToPLace int) maze {
	m := maze{}
	for i, input := range inputs {
		if i >= wallToPLace {
			break
		}
		m.addWall(input[0], input[1])
	}
	return m
}

func addBytesUntilImpossible(c []coords) coords {
	for i := 1024; i < len(c); i++ {
		m := createMaze(c, i)
		m.addSurroundingWalls()
		_, solved := m.solve()
		if !solved {
			return c[i-1]
		}
	}
	return coords{0, 0}
}

func main() {
	coords := readInput()
	m := createMaze(coords, 1024)
	m.addSurroundingWalls()
	pathLenth, _ := m.solve()
	fmt.Printf("Part 1 : %d\n", pathLenth)

	c := addBytesUntilImpossible(coords)
	fmt.Printf("Part 2 : %d,%d\n", c[0], c[1])
}
