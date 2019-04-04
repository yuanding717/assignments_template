package main

import (
  "fmt"
  // "sort"
  "strconv"
)

func CircularShiftRight(slice []int) {
  length := len(slice)
  last := slice[length - 1]
  for index := length - 1; index > 0; index-- {
    slice[index] = slice[index - 1]
  }
  slice[0] = last
}

func InsertToSorted(slice []int, newElement int) []int {
  if len(slice) == 0 {
    slice = append(slice, newElement)
  } else {
    slice = append(slice, newElement)
    indexOfFirstLarger := -1
    for index, element := range slice {
      if element > newElement {
        indexOfFirstLarger = index
        break
      }
    }
    if indexOfFirstLarger != -1 {
      CircularShiftRight(slice[indexOfFirstLarger : len(slice)])
    }
  }
  return slice
}

func main() {
  slice := make([]int, 0)
  var input string
  // count := 0
  for {
    fmt.Println("Please enter an integer or X to EXIT: ")
    if _, err := fmt.Scan(&input); err != nil {
      fmt.Println("Input error: ", err, ".Please try again.")
    } else {
      if (input == "X") {
        break
      }
      fmt.Println("Input string: ", input)
      if num, err := strconv.Atoi(input); err != nil {
        fmt.Println("Input string is not an integer number")
      } else {
        fmt.Println("Input string is an integer number: ", num)
        // if count < 3 {
        //   slice[count] = num
        // } else {
        //   slice = append(slice, num)
        // }
        // sorted := make([]int, len(slice))
        // copy(sorted, slice)
        // sort.Ints(sorted)
        // fmt.Println(slice)
        // fmt.Println(sorted)
        // count++
        slice = InsertToSorted(slice, num)
        fmt.Println("Slice in the (increasing) order: ", slice)
      }
    }
  }
}