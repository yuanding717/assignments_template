package main

import (
	"fmt"
)

func inputFloat64(name string) float64 {
  var input float64
  for {
    fmt.Printf("Please enter a float64 value for %s: ", name)
    if _, err := fmt.Scan(&input); err == nil {
      break
    } else {
      fmt.Println("Error: ", err)
    }
  }
  return input
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
  return func(time float64) float64 {
    return 0.5 * acceleration * time * time + initialVelocity * time + initialDisplacement
  }
}

func main() {
  acceleration := inputFloat64("acceleration")
  initialVelocity := inputFloat64("initialVelocity")
  initialDisplacement := inputFloat64("initialDisplacement")
  computeDisplacement := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
  time := inputFloat64("time")
  displacement := computeDisplacement(time)
  fmt.Printf("result (displacement) = %.2f\n", displacement)
}