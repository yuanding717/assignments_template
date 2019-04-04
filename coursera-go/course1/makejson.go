package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "strings"
)

type Person map[string]string

func main()  {
  person := make(Person)
  stdinReader := bufio.NewReader(os.Stdin)
  for {
		fmt.Print("Please enter the name: ")
		if name, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("Error: ", err)
		} else {
			person["name"] = strings.TrimRight(name, "\r\n")
			break
		}
  }
  for {
		fmt.Print("Please enter the address: ")
		if address, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("Error: ", err)
		} else {
			person["address"] = strings.TrimRight(address, "\r\n")
			break
		}
  }
  if data, err := json.Marshal(person); err != nil {
		fmt.Println("Error in serializing to JSON: ", err)
	} else {
		fmt.Println("JSON:", string(data))
	}
}