package main

import (
	"fmt"
	"os"
)

func readInput() []int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	dataInt := make([]int, 0)
	for _, c := range dataString {
		if c >= 0x30 && c <= 0x39 {
			dataInt = append(dataInt, int(c)-0x30)
		}
	}
	return dataInt
}

func uncompress(disk []int) []int {
	var withNumber int
	uncompressed := make([]int, 0)
	for i, d := range disk {
		toAppend := make([]int, d)
		if i%2 == 0 {
			withNumber = i / 2
		} else {
			withNumber = -1
		}
		for j := 0; j < len(toAppend); j++ {
			toAppend[j] = withNumber
		}
		uncompressed = append(uncompressed, toAppend...)
	}
	return uncompressed
}

func pop(disk []int) int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			poped := disk[i]
			disk[i] = -1
			return poped
		}
	}
	return -2
}

func finished(disk []int) bool {
	emptySeen := false
	for _, n := range disk {
		if n != -1 && emptySeen {
			return false
		}
		if n == -1 {
			emptySeen = true
		}
	}
	return true
}

func rearrange(disk []int) []int {
	arranged := make([]int, len(disk))
	copy(arranged, disk)
	for i, n := range arranged {
		if finished(arranged) {
			return arranged
		}
		if n == -1 {
			arranged[i] = pop(arranged)
		}
	}
	return arranged
}

func getfileSize(id int, originalDisk []int) int {
	return originalDisk[id*2]
}

func move(id int, offset int, size int, disk []int) {
	for i, n := range disk {
		if n == id {
			disk[i] = -1
		}
	}
	for i := 0; i < size; i++ {
		disk[i+offset] = id
	}
}

func findFreeSpace(id int, size int, disk []int) (int, bool) {
	offset := -1
	currentFreeSpace := 0
	started := false
	for i, n := range disk {
		if n == id {
			return offset, false
		}
		if n == -1 {
			if !started {
				offset = i
				started = true
			}
			currentFreeSpace++
			if currentFreeSpace >= size {
				return offset, true
			}
		} else {
			started = false
			currentFreeSpace = 0
		}
	}
	return offset, false
}

func newRearrange(originalDisk []int, disk []int) []int {
	var size, offset int
	var freeSpaceAvailable bool
	arranged := make([]int, len(disk))
	copy(arranged, disk)
	maxId := disk[len(disk)-1]
	for i := maxId; i > 0; i-- {
		size = getfileSize(i, originalDisk)
		offset, freeSpaceAvailable = findFreeSpace(i, size, arranged)
		if freeSpaceAvailable {
			move(i, offset, size, arranged)
		}
	}
	return arranged
}

func checksum(disk []int) int {
	sum := 0
	for i, n := range disk {
		if n != -1 {
			sum += n * i
		}
	}
	return sum
}

func main() {
	disk := readInput()
	uncompressed := uncompress(disk)
	rearranged := rearrange(uncompressed)
	fmt.Printf("Part 1 : %d\n", checksum(rearranged))

	newRearranged := newRearrange(disk, uncompressed)
	fmt.Printf("Part 2 : %d\n", checksum(newRearranged))
}
