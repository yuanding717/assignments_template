package main

import (
  "fmt"
)

func main()  {
  fmt.Printf("Please enter a floating point number: ")
  var f float32
  _, err := fmt.Scan(&f)
  if err != nil {
    fmt.Printf("Error: %s\n", err)
  } else {
    n := int32(f)
    fmt.Printf("Truncated value is %d\n", n)
  }
}