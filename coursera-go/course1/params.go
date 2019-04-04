package main

import (
  "fmt"
)

func modify(s []int) {
  fmt.Printf("%p \n", &s)
  s = []int{1,1,1,1}
  fmt.Println(s)
  fmt.Printf("%p \n", &s)
}

func modify2(s []int) {
  fmt.Printf("%p \n", &s)
  fmt.Printf("%p \n", &s[0])
  s[1] += 100
  fmt.Printf("%p \n", &s)
  fmt.Printf("%p \n", &s[0])
}

func main() {
  a := [5]int{1, 2, 3, 4, 5}
  s := a[:]
  // fmt.Printf("%p \n", &s)
  // modify(s)
  // fmt.Println(s[3])

  fmt.Printf("%p \n", &s)
  fmt.Printf("%p \n", &s[0])
  modify2(s)
  fmt.Println(a)
}