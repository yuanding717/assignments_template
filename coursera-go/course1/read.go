package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
  fname string
  lname string
}

func main() {
  var fileName string
  fmt.Println("Please enter the file name: ")
  if _, err := fmt.Scan(&fileName); err != nil {
    fmt.Println("Error: ", err)
    return
  } else {
    nameSlice := make([]name, 0, 20)
    fholder, err := os.Open(fileName)
    if err != nil {
      log.Fatalln(err)
      // os.Exit(1)
    }
    defer fholder.Close()
    sc := bufio.NewScanner(fholder)
    for sc.Scan() {
      line := sc.Text()
      words := strings.Fields(line)
      nameSlice = append(nameSlice, name{words[0], words[1]})
    }
    if err := sc.Err(); err != nil {
      log.Fatalf("scan file error: %v", err)
      return
    }
    for _, name := range nameSlice {
      fmt.Println(name.fname, name.lname)
    }
  }
}