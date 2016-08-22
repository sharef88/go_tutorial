package main

import (
	"fmt"
	"time"
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

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}
}
