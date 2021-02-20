package main

import (
	"fmt"
	"math"
)

//solution for Tour of Go excercise 8. Loops and Functions
//https://tour.golang.org/flowcontrol/8

func Sqrt(x float64) float64 {

	z := 1.0
	var oldz, diff float64

	for i := 1; i < 10 && (z > oldz || diff > .001); i++ {

		fmt.Printf("i = %v\n", i)
		fmt.Printf("oldz = %v\n", oldz)
		fmt.Printf("z = %v\n", z)
		fmt.Printf("diff: %v\n\n", diff)

		oldz = z
		z -= (z*z - x) / (2 * z) //Newton's method
		diff = oldz - z
	}
	return z
}

func main() {
	num := 5.0
	fmt.Printf("The squareroot of %v is %v\n", num, Sqrt(num))
	fmt.Printf("The math.Sqrt  of %v is %v\n\n", num, math.Sqrt(num))
}
