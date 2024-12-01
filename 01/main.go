package main

import (
	"fmt"
	"os"
  "strings"
  "strconv"
  "sort"
)

func readInput() ([]int, []int) {
  fileName := "input.txt"
  data, err := os.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
  dataString := string(data)
  linesNumber := strings.Count(dataString, "\n")
  list1, list2 := make([]int, linesNumber), make([]int, linesNumber)
  for i, line := range strings.Split(dataString, "\n") {
    if line == "" {
      continue
    }
    numbers := strings.Fields(line)
    num1, err1 := strconv.Atoi(numbers[0])
    num2, err2 := strconv.Atoi(numbers[1])
    list1[i], list2[i] = num1, num2
    if err1 != nil || err2 != nil {
			panic("Invalid number in input file")
		}
  }
  return list1, list2
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func similarity(x int, list []int) int {
  sim := 0
  for i := 0; i < len(list); i++ {
    if x == list[i] {
      sim++
    }
  }
  return sim
}

func main() {
  list1, list2 := readInput()
  sort.Ints(list1)
  sort.Ints(list2)
  sum := 0
  for i := 0; i < len(list1); i++ {
    sum += abs(list1[i] - list2[i])
  }
  fmt.Printf("Part 1 : %d\n", sum)

  sim := 0
  for i := 0; i < len(list1); i++ {
    sim += list1[i] * similarity(list1[i], list2)
  }
  fmt.Printf("Part 2 : %d\n", sim)
}
