package main

import (
  "fmt"
)

type Animal struct {
  locomotion string
	food       string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func getRequest() (string, string) {
	var name, info string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&name, &info); err != nil {
			fmt.Println("Error: ", err)
		} else {
			break
		}
	}
	return name, info
}

func main() {
  fmt.Println("Please type '[cow|bird|snake] [eat|move|speak]' after the prompt \">\" to discover animals. Press CTRL+C to exit.")
  animals := make(map[string]Animal)
	animals["cow"] = Animal{"walk", "grass", "moo"}
	animals["bird"] = Animal{locomotion: "fly", food: "worms", noise: "peep"}
  animals["snake"] = Animal{locomotion: "slither", food: "mice", noise: "hsss"}
  for {
		name, info := getRequest()
		animal, ok := animals[name]
		if !ok {
			fmt.Println("Unknown animal", name)
			continue
		}
		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("No info named '%s' about animal %s\n", info, name)
		}
	}
}