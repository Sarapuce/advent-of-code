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

func readInput() (map[string][][]int, []string) {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	input := strings.Split(dataString, "\n")
	input = removeEmptyStrings(input)
	antennas := make(map[string][][]int)
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[0]); y++ {
			antennaType := string(input[x][y])
			if antennaType != "." {
				_, ok := antennas[antennaType]
				if !ok {
					antennas[antennaType] = make([][]int, 0)
				}
				antennas[antennaType] = append(antennas[antennaType], []int{x, y})
			}
		}
	}
	m := make([]string, len(input))
	for i := range m {
		m[i] = strings.Repeat(".", len(input[0]))
	}
	return antennas, m
}

func getAntinode(x1, y1, x2, y2 int, m []string) [][]int {
	var a1, a2, b1, b2 int
	antinodes := make([][]int, 0)
	if x1 < x2 && y1 < y2 {
		a1, b1 = (2*x1)-x2, (2*y1)-y2
		a2, b2 = (2*x2)-x1, (2*y2)-y1
	} else if x1 >= x2 && y1 <= y2 {
		a1, b1 = (2*x1)-x2, (2*y1)-y2
		a2, b2 = (2*x2)-x1, (2*y2)-y1
	} else {
		return getAntinode(x2, y2, x1, y1, m)
	}
	if a1 >= 0 && a1 < len(m) && b1 >= 0 && b1 < len(m[0]) {
		antinodes = append(antinodes, []int{a1, b1})
	}
	if a2 >= 0 && a2 < len(m) && b2 >= 0 && b2 < len(m[0]) {
		antinodes = append(antinodes, []int{a2, b2})
	}
	return antinodes
}

func getNewAntinode(x1, y1, x2, y2 int, m []string) [][]int { // Probably something smart to do with negative number but I have to go play batllefield
	var a, b, a1, b1, a2, b2 int
	antinodes := make([][]int, 0)
	if x1 < x2 && y1 < y2 {
		a, b = x2-x1, y2-y1                     // a and b are the gap between antennas for x and y
		for i := 0; i < len(m)+len(m[0]); i++ { // Max len(m) len(m[0]) would have been better
			a1, b1 = x1+(i*a), y1+(i*b)
			a2, b2 = x1-(i*a), y1-(i*b)
			if a1 < len(m) && b1 < len(m[0]) {
				antinodes = append(antinodes, []int{a1, b1})
			}
			if a2 >= 0 && b2 >= 0 {
				antinodes = append(antinodes, []int{a2, b2})
			}
		}
	} else if x1 >= x2 && y1 <= y2 {
		a, b = x1-x2, y2-y1
		for i := 0; i < len(m)+len(m[0]); i++ {
			a1, b1 = x1-(i*a), y1+(i*b)
			a2, b2 = x1+(i*a), y1-(i*b)
			if a1 >= 0 && b1 < len(m[0]) {
				antinodes = append(antinodes, []int{a1, b1})
			}
			if a2 < len(m) && b2 >= 0 {
				antinodes = append(antinodes, []int{a2, b2})
			}
		}
	} else {
		return getNewAntinode(x2, y2, x1, y1, m)
	}
	return antinodes
}

func processAntenna(coordinates [][]int, m []string) [][]int {
	antinodes := make([][]int, 0)
	for i, coordinate1 := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			coordinate2 := coordinates[j]
			antinodes = append(antinodes, getAntinode(coordinate1[0], coordinate1[1], coordinate2[0], coordinate2[1], m)...)
		}
	}
	return antinodes
}

func processNewAntenna(coordinates [][]int, m []string) [][]int {
	antinodes := make([][]int, 0)
	for i, coordinate1 := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			coordinate2 := coordinates[j]
			fmt.Println(coordinate1)
			fmt.Println(coordinate2)
			fmt.Println()
			antinodes = append(antinodes, getNewAntinode(coordinate1[0], coordinate1[1], coordinate2[0], coordinate2[1], m)...)
		}
	}
	antinodes = append(antinodes, coordinates...)
	return antinodes
}

func sumAntinode(antinodes [][]int) int {
	sum := 0
	alreadySeen := make([][]int, 0)
	seen := false
	for _, antinode := range antinodes {
		seen = false
		for _, seenAntinode := range alreadySeen {
			if antinode[0] == seenAntinode[0] && antinode[1] == seenAntinode[1] {
				seen = true
				break
			}
		}
		if !seen {
			sum += 1
			alreadySeen = append(alreadySeen, antinode)
		}
	}
	return sum
}

func printAntenna(antinodes [][]int, coordinates [][]int, m []string) {
	x := len(m)
	y := len(m[0])
	antennaMap := make([]string, x)
	for i := range y {
		antennaMap[i] = strings.Repeat(".", y)
	}
	for _, c := range coordinates {
		antennaMap[c[0]] = antennaMap[c[0]][:c[1]] + "x" + antennaMap[c[0]][c[1]+1:]
	}
	for _, a := range antinodes {
		antennaMap[a[0]] = antennaMap[a[0]][:a[1]] + "#" + antennaMap[a[0]][a[1]+1:]
	}
	for _, line := range antennaMap {
		fmt.Println(line)
	}
	fmt.Println()
}

func main() {
	antennas, m := readInput()
	antinodes := make([][]int, 0)
	for _, antenna := range antennas {
		antinodes = append(antinodes, processAntenna(antenna, m)...)
	}
	sum := sumAntinode(antinodes)
	fmt.Printf("Part 1 : %d\n", sum)

	newAntinodes := make([][]int, 0)
	for _, antenna := range antennas {
		newAntinodes = append(newAntinodes, processNewAntenna(antenna, m)...)
		printAntenna(newAntinodes, antenna, m)
	}
	sum = sumAntinode(newAntinodes)
	fmt.Printf("Part 2 : %d\n", sum)
}
