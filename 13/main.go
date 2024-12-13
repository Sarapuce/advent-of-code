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

func readInput() [][3][2]int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	inputs := strings.Split(dataString, "\n\n")
	inputs = removeEmptyStrings(inputs)

	var input, numberString string
	var lines []string
	var Ax, Ay, Bx, By, Mx, My int
	machines := make([][3][2]int, len(inputs))
	for i := 0; i < len(machines); i++ {
		input = inputs[i]
		lines = strings.Split(input, "\n")
		numberString = strings.Split(strings.Split(lines[0], ",")[0], "+")[1]
		Ax, _ = strconv.Atoi(numberString)
		numberString = strings.Split(strings.Split(lines[0], ",")[1], "+")[1]
		Ay, _ = strconv.Atoi(numberString)

		numberString = strings.Split(strings.Split(lines[1], ",")[0], "+")[1]
		Bx, _ = strconv.Atoi(numberString)
		numberString = strings.Split(strings.Split(lines[1], ",")[1], "+")[1]
		By, _ = strconv.Atoi(numberString)

		numberString = strings.Split(strings.Split(lines[2], ",")[0], "=")[1]
		Mx, _ = strconv.Atoi(numberString)
		numberString = strings.Split(strings.Split(lines[2], ",")[1], "=")[1]
		My, _ = strconv.Atoi(numberString)

		machines[i][0][0], machines[i][0][1] = Ax, Ay
		machines[i][1][0], machines[i][1][1] = Bx, By
		machines[i][2][0], machines[i][2][1] = Mx, My
	}
	return machines
}

func findMax(target [2]int, b [2]int) int {
	var maxX, maxY int
	maxX = target[0] / b[0]
	maxY = target[1] / b[1]
	if maxX > maxY {
		if maxY > 100 {
			return 100
		} else {
			return maxY
		}
	} else if maxX > 100 {
		return 100
	} else {
		return maxX
	}
}

func brutforceOptimumB(machine [3][2]int) (int, int) {
	var gapX, gapY, aPress int
	a, b, target := machine[0], machine[1], machine[2]
	maxB := findMax(target, b)
	for bPress := maxB; bPress >= 0; bPress-- {
		gapX = target[0] - (b[0] * bPress)
		if gapX%a[0] == 0 {
			aPress = gapX / a[0]
			gapY = target[1] - (b[1] * bPress)
			if gapY%a[1] == 0 && gapY/a[1] == aPress && aPress <= 100 {
				return aPress, bPress
			}
		}
	}
	return -1, -1
}

func brutforceOptimumA(machine [3][2]int) (int, int) {
	var gapX, gapY, bPress int
	a, b, target := machine[0], machine[1], machine[2]
	maxA := findMax(target, a)
	for aPress := maxA; aPress >= 0; aPress-- {
		gapX = target[0] - (a[0] * aPress)
		if gapX%b[0] == 0 {
			bPress = gapX / b[0]
			gapY = target[1] - (a[1] * aPress)
			if gapY%b[1] == 0 && gapY/b[1] == bPress && bPress <= 100 {
				return aPress, bPress
			}
		}
	}
	return -1, -1
}

func calcPrice(machines [][3][2]int) int {
	totalPrice := 0
	var a, b, price int
	for _, machine := range machines {
		a, b = brutforceOptimumB(machine)
		if a == -1 {
			price = 0
		} else {
			price = a*3 + b
		}
		a, b = brutforceOptimumA(machine)
		if a != -1 && (a*3+b < price || price == 0) {
			price = a*3 + b
		}
		totalPrice += price
	}
	return totalPrice
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getFirstApprox(machine [3][2]int) (int, int) {
	landing := [2]int{0, 0}
	minimum := 2000
	minA, minB := 0, 0
	var distance int
	a, b, target := machine[0], machine[1], machine[2]
	target[0], target[1] = 1000, 1000
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			landing = [2]int{i*a[0] + j*b[0], i*a[1] + j*b[1]}
			distance = abs(landing[0]-target[0]) + abs(landing[1]-target[1])
			if distance < minimum {
				minimum = distance
				minA = i
				minB = j
			}
		}
	}
	return minA * 10000000000, minB * 10000000000
}

func getCloseTo(machine [3][2]int, start [2]int, scale int) (int, int) {
	landing := [2]int{0, 0}
	minimum := 2000000
	startA, startB := start[0]/scale, start[1]/scale
	minA, minB := 0, 0
	var distance int
	a, b, target := machine[0], machine[1], machine[2]
	target[0], target[1] = target[0]+10000000000000, target[1]+10000000000000
	target[0], target[1] = target[0]/scale, target[1]/scale
	for i := -1000; i < 1000; i++ {
		for j := -1000; j < 1000; j++ {
			landing = [2]int{(i+startA)*a[0] + (j+startB)*b[0], (i+startA)*a[1] + (j+startB)*b[1]}
			distance = abs(landing[0]-target[0]) + abs(landing[1]-target[1])
			if distance < minimum {
				minimum = distance
				minA = i
				minB = j
			}
		}
	}
	return (startA + minA) * scale, (startB + minB) * scale
}

func calcNewPrice(machines [][3][2]int) int {
	priceTotal := 0
	var approxA, approxB int
	for _, machine := range machines {
		a, b, target := machine[0], machine[1], machine[2]
		target[0], target[1] = target[0]+10000000000000, target[1]+10000000000000
		approxA, approxB = getFirstApprox(machine)
		for scale := 1000000000; scale >= 1; scale /= 10 {
			approxA, approxB = getCloseTo(machine, [2]int{approxA, approxB}, scale)
		}
		if (approxA*a[0]+approxB*b[0] == target[0]) && (approxA*a[1]+approxB*b[1] == target[1]) {
			priceTotal += 3*approxA + approxB
		}
	}
	return priceTotal
}

func main() {
	machines := readInput()
	price := calcPrice(machines)
	fmt.Printf("Part 1 : %d\n", price)

	price = calcNewPrice(machines)
	fmt.Printf("Part 2 : %d\n", price)
}
