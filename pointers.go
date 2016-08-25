package main

import "fmt"

type Vertex struct {
	//its a vertex declaration? x,y
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{} //as all values default to their 0 or false value, undef -> 0
	p  = &Vertex{1, 2}
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

}
