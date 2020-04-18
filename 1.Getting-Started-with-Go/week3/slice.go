package main

import (
	"fmt"
	"sort"
)

// Numbers implements the interface required by sort.Sort()
type Numbers []int

func (s Numbers) Len() int {
	return len(s)
}

func (s Numbers) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Numbers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := make([]int, 0)
	for {
		fmt.Println("Please enter an integer:")
		var x int
		nRead, err := fmt.Scan(&x)
		if nRead != 1 || err != nil {
			break
		}
		s = append(s, x)
		sort.Sort(Numbers(s))
		fmt.Println(s)
	}
}
