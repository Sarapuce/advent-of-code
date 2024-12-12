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

func findNext(alreadyDone [][]bool) (int, int) {
	for x := 0; x < len(alreadyDone); x++ {
		for y := 0; y < len(alreadyDone[0]); y++ {
			if !alreadyDone[x][y] {
				return x, y
			}
		}
	}
	return -1, -1
}

func findNeighbor(garden []string, x int, y int, plant string, area [][]bool) [][2]int {
	neighbor := make([][2]int, 0)
	if x+1 < len(garden) {
		if string(garden[x+1][y]) == plant && !area[x+1][y] {
			area[x+1][y] = true
			neighbor = append(neighbor, [2]int{x + 1, y})
		}
	}
	if x-1 >= 0 {
		if string(garden[x-1][y]) == plant && !area[x-1][y] {
			area[x-1][y] = true
			neighbor = append(neighbor, [2]int{x - 1, y})
		}
	}
	if y+1 < len(garden[0]) {
		if string(garden[x][y+1]) == plant && !area[x][y+1] {
			area[x][y+1] = true
			neighbor = append(neighbor, [2]int{x, y + 1})
		}
	}
	if y-1 >= 0 {
		if string(garden[x][y-1]) == plant && !area[x][y-1] {
			area[x][y-1] = true
			neighbor = append(neighbor, [2]int{x, y - 1})
		}
	}
	return neighbor
}

func getArea(garden []string, startX int, startY int) [][]bool {
	area := make([][]bool, len(garden))
	for i := 0; i < len(garden); i++ {
		area[i] = make([]bool, len(garden[0]))
	}
	neighbor := make([][2]int, 1)
	neighbor[0] = [2]int{startX, startY}
	plant := string(garden[startX][startY])
	area[startX][startY] = true
	for len(neighbor) > 0 {
		newNeighbor := make([][2]int, 0)
		for _, n := range neighbor {
			newNeighbor = append(newNeighbor, findNeighbor(garden, n[0], n[1], plant, area)...)
		}
		neighbor = newNeighbor
	}
	return area
}

func updateDone(alreadyDone [][]bool, area [][]bool) {
	for x := 0; x < len(alreadyDone); x++ {
		for y := 0; y < len(alreadyDone[0]); y++ {
			alreadyDone[x][y] = alreadyDone[x][y] || area[x][y]
		}
	}
}

func printArea(area [][]bool) {
	s := ""
	for x := 0; x < len(area); x++ {
		for y := 0; y < len(area[0]); y++ {
			if area[x][y] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}

func getPerimeter(area [][]bool) int {
	perimeter := 0
	oldY, oldX := false, false
	for x := 0; x < len(area); x++ {
		oldX = false
		for y := 0; y < len(area[0]); y++ {
			if oldX != area[x][y] {
				perimeter += 1
				oldX = area[x][y]
			}
		}
		if oldX {
			perimeter += 1
		}
	}
	for y := 0; y < len(area); y++ {
		oldY = false
		for x := 0; x < len(area[0]); x++ {
			if oldY != area[x][y] {
				perimeter += 1
				oldY = area[x][y]
			}
		}
		if oldY {
			perimeter += 1
		}
	}
	return perimeter
}

// Check if 2 squares side to side are different. If so, change the state of xorX and add a part of perimeter each time it is done
func getNewPerimeter(area [][]bool) int {
	perimeter := 0
	wideArea := make([][]bool, len(area)+2)
	for i := 0; i < len(wideArea); i++ {
		wideArea[i] = make([]bool, len(area[0])+2)
	}
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {
			wideArea[i+1][j+1] = area[i][j]
		}
	}
	oldY, oldX := false, false
	xorX, xorY := false, false
	currentPerimeter := 0
	for y := 0; y < len(wideArea[0])-1; y++ {
		oldX = false
		currentPerimeter = 0
		for x := 0; x < len(wideArea); x++ {
			xorX = wideArea[x][y] != wideArea[x][y+1]
			if oldX != xorX {
				currentPerimeter += 1
				oldX = xorX
			} else if xorX { // Handle the special case where you have a corner inside the area like this :
				// ####
				// #.##
				// ##.#
				// ####
				if wideArea[x][y] != wideArea[x-1][y] { // Safe because of the wideArea
					currentPerimeter += 2 // Add 2 because we will divide by 2
				}
			}
		}
		perimeter += currentPerimeter / 2
		// Because xorX switch to 1 and to 0 each time a perimeter is found, you have to divide by 2
	}
	for x := 0; x < len(wideArea)-1; x++ {
		oldY = false
		currentPerimeter = 0
		for y := 0; y < len(wideArea[0]); y++ {
			xorY = wideArea[x][y] != wideArea[x+1][y]
			if oldY != xorY {
				currentPerimeter += 1
				oldY = xorY
			} else if xorY {
				if wideArea[x][y] != wideArea[x][y-1] {
					currentPerimeter += 2
				}
			}
		}
		perimeter += currentPerimeter / 2
	}
	return perimeter
}

func getSurface(area [][]bool) int {
	surface := 0
	for x := 0; x < len(area); x++ {
		for y := 0; y < len(area[0]); y++ {
			if area[x][y] {
				surface += 1
			}
		}
	}
	return surface
}

func main() {
	garden := readInput()
	price := 0
	alreadyDone := make([][]bool, len(garden))
	for i := 0; i < len(garden); i++ {
		alreadyDone[i] = make([]bool, len(garden[0]))
	}
	for startX, startY := findNext(alreadyDone); startX != -1; startX, startY = findNext(alreadyDone) {
		area := getArea(garden, startX, startY)
		perimeter := getPerimeter(area)
		surface := getSurface(area)
		price += perimeter * surface
		updateDone(alreadyDone, area)
	}

	fmt.Printf("Part 1 : %d\n", price)

	price = 0
	alreadyDone = make([][]bool, len(garden))
	for i := 0; i < len(garden); i++ {
		alreadyDone[i] = make([]bool, len(garden[0]))
	}
	for startX, startY := findNext(alreadyDone); startX != -1; startX, startY = findNext(alreadyDone) {
		area := getArea(garden, startX, startY)
		perimeter := getNewPerimeter(area)
		surface := getSurface(area)
		price += perimeter * surface
		updateDone(alreadyDone, area)
	}
	fmt.Printf("Part 2 : %d\n", price)
}
