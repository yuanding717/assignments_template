package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  slice := ReadSlice()
  BubbleSort(slice)
  fmt.Println(slice)
}

func ReadSlice() []int {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Please enter a size of array from 0 to 10: ")
  scanner.Scan()
  size, err := strconv.Atoi(scanner.Text())
  if err != nil || size < 0 || size > 10 {
    log.Fatalln("Size is not correct")
  }
  fmt.Println("Please enter an array to sort (press enter after each element of the array): ")
  result := make([]int, size)
  for i := 0; i < size; i++ {
    scanner.Scan()
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln("Passed incorrect integer")
		}
		result[i] = item
  }
  return result
}

func BubbleSort(slice []int) {
  for i := 0; i < len(slice); i++ {
    for j := i + 1; j < len(slice); j++ {
      if (slice[i] > slice[j]) {
        Swap(slice, i, j)
      }
    }
  }
}

func Swap(slice []int, i, j int) {
  slice[i], slice[j] = slice[j], slice[i]
}