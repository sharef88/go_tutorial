package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	//z = guess
	z := float64(100)

	//Newton's method for finding the sqrt.  This is mildly obnoxious
	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}

	//step through the equation count times
	count := 10
	for index := 0; index < count; index++ {
		z = step()
	}
	return z

}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(10), math.Sqrt(10))
}
