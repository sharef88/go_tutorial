package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	var s []int
	printSlice("s", s)

	//append works on nil slices
	s = append(s, 0)
	printSlice("s", s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice("s", s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice("s", s)

}
