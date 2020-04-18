package main

import "fmt"

// Swap s[i] and s[i + 1]
func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

// BubbleSort sorts a slice using the bubble sort algorithm
func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
}

func main() {

	s := []int{}

	fmt.Println("Please enter some integers (10 integers at most) separated by space (enter x to end the input):")
	for i := 0; i < 10; i++ {
		var x int
		nRead, err := fmt.Scan(&x)
		if nRead != 1 || err != nil {
			break
		}
		s = append(s, x)
	}
	fmt.Println("\nOriginal slice:")
	fmt.Println(s)
	BubbleSort(s)
	fmt.Println("\nSorted slice:")
	fmt.Println(s)
}
