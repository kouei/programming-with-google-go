package main

import (
	"fmt"
	"os"
)

func readFloat() float64 {
	var f float64
	nRead, err := fmt.Scan(&f)
	if nRead != 1 || err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return f
}

// GenDisplaceFn returns a function calculate displacement according to time
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	fmt.Println("Please enter acceleration:")
	a := readFloat()

	fmt.Println("Please enter initial speed:")
	v0 := readFloat()

	fmt.Println("Please enter initial displacement:")
	s0 := readFloat()

	getDisplace := GenDisplaceFn(a, v0, s0)

	fmt.Println("Please enter time:")
	t := readFloat()

	fmt.Println("\nThe displacement is:", getDisplace(t))
}
