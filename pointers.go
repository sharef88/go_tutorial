package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	ry := make([][]uint8, dy)
	for i := range ry {
		ry[i] = make([]uint8, dx)
		for j := range ry[i] {
			val := (j ^ i)
			ry[i][j] = uint8(val)
		}
	}
	return ry
}

func main() {
	pic.Show(Pic)
}
