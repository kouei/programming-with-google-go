// Race-condition means the output of a program depends on the order of execution of its threads.
// This program is an example to illustrate race-condition using goroutines.
package main

import (
	"fmt"
	"sync"
)

func f(x *int, wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		*x++
	}
	defer wg.Done()
}

func main() {
	x := 0

	wg := sync.WaitGroup{}

	// Here, two goroutines will run alternately.
	// Each of them is trying to add 100000 to x.
	// The final value of x depends on the order of (instruction level) execution of these two goroutines.
	wg.Add(2)
	go f(&x, &wg)
	go f(&x, &wg)
	wg.Wait()

	// the output will be a random number in range [100000, 200000].
	fmt.Println(x)
}
