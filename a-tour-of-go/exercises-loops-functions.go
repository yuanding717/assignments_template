package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// initial guess
	z := 1.0
	diff := math.Abs(x - z)
	for (diff > 1e-6) {
		new_z := z - (z*z - x) / (2*z)
		diff = math.Abs(z - new_z)
		z = new_z
	}
	return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(3))
}