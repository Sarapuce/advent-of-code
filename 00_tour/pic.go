package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dx)
	fmt.Println(p)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			p[y] = append(p[y], uint8((x+y)/2))
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
