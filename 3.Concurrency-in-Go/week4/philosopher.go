// This program prints different log information from the specification of the assignment.
// The message printed by this program is more informative.

package main

import (
	"fmt"
	"sync"
)

// Philosopher struct
type Philosopher struct {
	id int
}

// Chopstick struct
type Chopstick struct {
	mut sync.Mutex
}

// Host struct
type Host struct {
	givePermission   chan int
	returnPermission chan int
}

// Eat method of Philosopher
func (p Philosopher) Eat(host Host, cs []Chopstick, i int) {

	n := len(cs)
	suffix := map[int]string{
		1: "st",
		2: "nd",
		3: "rd",
	}

	<-host.givePermission // get permission from the host
	defer func() { host.returnPermission <- 1 }()

	fmt.Printf("[Philosopher %d] [%d%s] eat started\n", p.id, i, suffix[i])

	cs[p.id-1].mut.Lock() // get left chopstick
	defer cs[p.id-1].mut.Unlock()

	cs[p.id%n].mut.Lock() // get right chopstick
	defer cs[p.id%n].mut.Unlock()

	fmt.Printf("[Philosopher %d] [%d%s] eat finished\n", p.id, i, suffix[i])

}

// Run method of Philosopher
func (p Philosopher) Run(host Host, cs []Chopstick, wg *sync.WaitGroup, finished chan int) {
	defer wg.Done()
	defer func() { finished <- 1 }()

	nEat := 3
	for i := 1; i <= nEat; i++ {
		p.Eat(host, cs, i)
	}
}

// Run method of Host
func (host Host) Run(wg *sync.WaitGroup, finished chan int) {
	defer wg.Done()

	count := 0
	for count < cap(finished) {
		select {
		case <-finished: // check if there is any philosopher finished its eat
			count++
		case <-host.returnPermission: // check if it is possible to release another permission
			host.givePermission <- 1
		}
	}
	fmt.Printf("[Host] finished its job\n")
}

func main() {

	wg := sync.WaitGroup{}

	nPhilosopher := 5
	ps := make([]Philosopher, nPhilosopher)
	for i := range ps {
		ps[i].id = i + 1
	}

	cs := make([]Chopstick, nPhilosopher)

	nPermission := 2
	host := Host{make(chan int, nPermission), make(chan int, nPermission)} // buffered channel of size 2, so that there are always up to 2 permissions
	for i := 0; i < nPermission; i++ {
		host.returnPermission <- 1
	}

	finished := make(chan int, nPhilosopher)

	wg.Add(1)
	go host.Run(&wg, finished)

	for _, p := range ps {
		wg.Add(1)
		go p.Run(host, cs, &wg, finished)
	}
	wg.Wait()
}
