//+build disabled

package main

import (
	"fmt"
	"time"
)

var global int = 10

func checkGlobal() {
	if global == 10 {
		fmt.Println("global is 10")
	} else {
		fmt.Println("global is not 10!")
	}
}

func incrementGlobal() {
	global++
}

func main() {
	go checkGlobal()
	go incrementGlobal()

	/* sleep for a bit so that the results are actually printed out */
	time.Sleep(200 * time.Millisecond)
}
