package main

import (
  "strings"
  "fmt"
  "bufio"
  "os"
)

func hasIAN(input string) bool {
  sanitized := strings.TrimSpace(strings.ToLower(input))
  if (strings.HasPrefix(sanitized, "i") && strings.Contains(sanitized, "a") && strings.HasSuffix(sanitized, "n")) {
    return true
  }
  return false
}

func main () {
  // var input string
  fmt.Println("Please enter a string: ")
  // _, err := fmt.Scan(&input)
  // if err != nil {
  //   fmt.Printf("Error: %s\n", err)
  // } else {
  //   if (hasIAN(input)) {
  //     fmt.Println("Found!")
  //   } else {
  //     fmt.Println("Not Found!")
  //   }
  // }
  stdinReader := bufio.NewReader(os.Stdin)
  if input, err := stdinReader.ReadString('\n'); err != nil {
    fmt.Printf("Error: %s\n", err)
  } else {
    if (hasIAN(input)) {
      fmt.Println("Found!")
    } else {
      fmt.Println("Not Found!")
    }
  }
}