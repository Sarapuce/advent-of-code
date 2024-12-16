package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() ([][]string, string) {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	inputs := strings.Split(dataString, "\n\n")

	area := inputs[0]
	areaLine := strings.Split(area, "\n")
	m := make([][]string, len(areaLine))
	for i, line := range areaLine {
		m[i] = make([]string, len(line))
		for j, c := range line {
			m[i][j] = string(c)
		}
	}

	commands := inputs[1]
	commands = strings.ReplaceAll(commands, "\n", "")

	return m, commands
}

func getNewMap(m [][]string) [][]string {
	newMap := make([][]string, len(m))
	for i := 0; i < len(newMap); i++ {
		newMap[i] = make([]string, len(m[0])*2)
	}
	for i, line := range m {
		for j, object := range line {
			if object == "." {
				newMap[i][j*2], newMap[i][(j*2)+1] = ".", "."
			} else if object == "#" {
				newMap[i][j*2], newMap[i][(j*2)+1] = "#", "#"
			} else if object == "@" {
				newMap[i][j*2], newMap[i][(j*2)+1] = "@", "."
			} else if object == "O" {
				newMap[i][j*2], newMap[i][(j*2)+1] = "[", "]"
			} else {
				fmt.Println("Can't exec getNewMap")
			}
		}
	}
	return newMap
}

func printMap(m [][]string) {
	for _, line := range m {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func getRobotPos(m [][]string) (int, int) {
	for x, line := range m {
		for y, c := range line {
			if c == "@" {
				return x, y
			}
		}
	}
	return -1, -1
}

func canPushBox(m [][]string, command string, x int, y int) bool {
	var newX, newY int
	switch command {
	case "v":
		newX, newY = x+1, y
	case "^":
		newX, newY = x-1, y
	case ">":
		newX, newY = x, y+1
	case "<":
		newX, newY = x, y-1
	default:
		fmt.Println("Can't exec canPushBox")
	}

	if m[newX][newY] == "#" {
		return false
	} else if m[newX][newY] == "." {
		return true
	} else if m[newX][newY] == "O" {
		return canPushBox(m, command, newX, newY)
	}
	fmt.Println("Can't exec canPushBox")
	return false
}

func canPushNewBox(m [][]string, command string, x int, y int, boxPart string) bool {
	var newX, newY int
	mTemp := make([][]string, len(m))
	for i := range m {
		mTemp[i] = make([]string, len(m[i]))
		copy(mTemp[i], m[i])
	}
	switch command {
	case "v":
		newX, newY = x+1, y
	case "^":
		newX, newY = x-1, y
	case ">":
		newX, newY = x, y+1
	case "<":
		newX, newY = x, y-1
	default:
		fmt.Println("Can't exec canPushNewBox")
	}

	if boxPart == "[" {
		mTemp[x][y] = "."
		mTemp[x][y+1] = "."
		if m[newX][newY] == "#" || m[newX][newY+1] == "#" {
			return false
		} else if mTemp[newX][newY] == "." && mTemp[newX][newY+1] == "." {
			return true
		} else if mTemp[newX][newY] == "[" {
			return canPushNewBox(m, command, newX, newY, "[")
		} else if mTemp[newX][newY] == "]" && mTemp[newX][newY+1] == "." {
			return canPushNewBox(m, command, newX, newY, "]")
		} else if mTemp[newX][newY+1] == "[" && mTemp[newX][newY] == "." {
			return canPushNewBox(m, command, newX, newY+1, "[")
		} else if mTemp[newX][newY] == "]" && mTemp[newX][newY+1] == "[" {
			return canPushNewBox(m, command, newX, newY, "]") && canPushNewBox(m, command, newX, newY+1, "[")
		} else {
			fmt.Println("Can't execute canPushNewBox")
		}
	} else if boxPart == "]" {
		mTemp[x][y] = "."
		mTemp[x][y-1] = "."
		if m[newX][newY] == "#" || m[newX][newY-1] == "#" {
			return false
		} else if mTemp[newX][newY] == "." && mTemp[newX][newY-1] == "." {
			return true
		} else if mTemp[newX][newY] == "]" {
			return canPushNewBox(m, command, newX, newY, "]")
		} else if mTemp[newX][newY] == "[" && mTemp[newX][newY-1] == "." {
			return canPushNewBox(m, command, newX, newY, "[")
		} else if mTemp[newX][newY-1] == "]" && mTemp[newX][newY] == "." {
			return canPushNewBox(m, command, newX, newY-1, "]")
		} else if mTemp[newX][newY] == "[" && mTemp[newX][newY-1] == "]" {
			return canPushNewBox(m, command, newX, newY, "[") && canPushNewBox(m, command, newX, newY-1, "]")
		} else {
			fmt.Println("Can't execute canPushNewBox")
		}
	}
	return false
}

func pushBox(m [][]string, command string, x int, y int) {
	var newX, newY int
	switch command {
	case "v":
		newX, newY = x+1, y
	case "^":
		newX, newY = x-1, y
	case ">":
		newX, newY = x, y+1
	case "<":
		newX, newY = x, y-1
	default:
		fmt.Println("Can't exec pushBox")
	}

	if m[newX][newY] == "." {
		m[newX][newY] = "O"
	} else if m[newX][newY] == "O" {
		pushBox(m, command, newX, newY)
	}
}

func pushNewBox(m [][]string, command string, x int, y int, boxPart string) {
	var newX, newY int
	mTemp := make([][]string, len(m))
	for i := range m {
		mTemp[i] = make([]string, len(m[i]))
		copy(mTemp[i], m[i])
	}
	switch command {
	case "v":
		newX, newY = x+1, y
	case "^":
		newX, newY = x-1, y
	case ">":
		newX, newY = x, y+1
	case "<":
		newX, newY = x, y-1
	default:
		fmt.Println("Can't exec pushBox")
	}

	if boxPart == "[" {
		mTemp[x][y] = "."
		mTemp[x][y+1] = "." // Get a map status without the current box
		if mTemp[newX][newY] == "[" || mTemp[newX][newY] == "]" {
			pushNewBox(m, command, newX, newY, mTemp[newX][newY])
			for i := range m {
				mTemp[i] = make([]string, len(m[i]))
				copy(mTemp[i], m[i])
			}
			mTemp[x][y] = "."
			mTemp[x][y+1] = "." // mTemp must be sync with m (sorry it's late)
		}
		if mTemp[newX][newY+1] == "[" || mTemp[newX][newY+1] == "]" {
			pushNewBox(m, command, newX, newY+1, mTemp[newX][newY+1])
		}
		m[x][y], m[x][y+1] = ".", "."
		m[newX][newY], m[newX][newY+1] = "[", "]"
	} else if boxPart == "]" {
		mTemp[x][y] = "."
		mTemp[x][y-1] = "." // Can be factorized, but it's late
		if mTemp[newX][newY] == "[" || mTemp[newX][newY] == "]" {
			pushNewBox(m, command, newX, newY, mTemp[newX][newY])
			for i := range m {
				mTemp[i] = make([]string, len(m[i]))
				copy(mTemp[i], m[i])
			}
			mTemp[x][y] = "."
			mTemp[x][y+1] = "."
		}
		if mTemp[newX][newY-1] == "[" || mTemp[newX][newY-1] == "]" {
			pushNewBox(m, command, newX, newY-1, mTemp[newX][newY-1])
		}
		m[x][y], m[x][y-1] = ".", "."
		m[newX][newY-1], m[newX][newY] = "[", "]"
	}
}

func moveRobot(m [][]string, x int, y int, newX int, newY int) {
	m[newX][newY] = "@"
	m[x][y] = "."
}

func move(m [][]string, command string) {
	x, y := getRobotPos(m)
	var newX, newY int
	switch command {
	case "v":
		newX, newY = x+1, y
	case "^":
		newX, newY = x-1, y
	case ">":
		newX, newY = x, y+1
	case "<":
		newX, newY = x, y-1
	default:
		fmt.Println("Can't exec move")
	}

	if m[newX][newY] == "." {
		moveRobot(m, x, y, newX, newY)
	} else if m[newX][newY] == "O" {
		if canPushBox(m, command, newX, newY) {
			pushBox(m, command, newX, newY)
			moveRobot(m, x, y, newX, newY)
		}
	}
}

func newMove(m [][]string, command string) {
	x, y := getRobotPos(m)
	var newX, newY int
	switch command {
	case "v":
		newX, newY = x+1, y
	case "^":
		newX, newY = x-1, y
	case ">":
		newX, newY = x, y+1
	case "<":
		newX, newY = x, y-1
	default:
		fmt.Println("Can't exec move")
	}

	if m[newX][newY] == "." {
		moveRobot(m, x, y, newX, newY)
	} else if m[newX][newY] == "[" {
		if canPushNewBox(m, command, newX, newY, "[") {
			pushNewBox(m, command, newX, newY, "[")
			moveRobot(m, x, y, newX, newY)
		}
	} else if m[newX][newY] == "]" {
		if canPushNewBox(m, command, newX, newY, "]") {
			pushNewBox(m, command, newX, newY, "]")
			moveRobot(m, x, y, newX, newY)
		}
	}
}

func checksum(m [][]string) int {
	sum := 0
	for x, line := range m {
		for y, object := range line {
			if object == "O" {
				sum += (x * 100) + y
			}
		}
	}
	return sum
}

func newChecksum(m [][]string) int {
	sum := 0
	for x, line := range m {
		for y, object := range line {
			if object == "[" {
				sum += (x * 100) + y
			}
		}
	}
	return sum
}

func main() {
	m, commands := readInput()
	for _, c := range commands {
		move(m, string(c))
	}
	fmt.Printf("Part 1 : %d\n", checksum(m))

	m, commands = readInput()
	newMap := getNewMap(m)
	for _, c := range commands {
		newMove(newMap, string(c))
	}
	fmt.Printf("Part 2 : %d\n", newChecksum(newMap))
}
