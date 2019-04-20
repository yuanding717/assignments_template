package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

// Used as semaphore with value of 2.
// That allows no more than 2 philosophers to eat concurrently.
var sem = make(chan int, 2)

var wg sync.WaitGroup

func (p Philo) eat() {
	defer wg.Done()

	// Eat only 3 times.
	for {
		sem <- 1 // acquire permission
		chops := getChopsInRandom(p)
		takeOffChops(chops)

		// Eating.
		fmt.Printf("starting to eat %d\n", p.number)
		//time.Sleep(20 * time.Millisecond) // waiting to demonstrate concurrency
		fmt.Printf("finishing eating %d\n", p.number)

		putBackChops(chops)
		<-sem // release permission
	}
}

// Put chopsticks on the table.
func putBackChops(chops [2]*ChopS) {
	for _, chop := range chops {
		chop.Unlock()
	}
}

// Get chopsticks off the table.
func takeOffChops(chops [2]*ChopS) {
	for _, chop := range chops {
		chop.Lock()
	}
}

// Randomize chopsticks.
func getChopsInRandom(p Philo) [2]*ChopS {
	chops := [2]*ChopS{p.leftCS, p.rightCS}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(2, func(i, j int) {
		chops[i], chops[j] = chops[j], chops[i]
	})
	return chops
}

func main() {
	// Initialize chopsticks.
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Initialize philosophers.
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	wg.Add(5)

	// Start eating.
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
