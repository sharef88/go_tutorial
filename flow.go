package main

import "fmt"

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

func main() {
	fmt.Println("counting")

	for index := 0; index < 10; index++ {
		defer fmt.Println(index)
	}
	fmt.Println("done")
}
