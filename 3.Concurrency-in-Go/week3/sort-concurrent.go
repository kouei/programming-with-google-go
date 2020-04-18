package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortLinear(s []int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("\nGoroutine-%v is sorting %v\n", id, s)
	sort.Ints(s)
}

// merge many slices into a single slice
func merge(ss [][]int) []int {
	res := []int{}
	index := []int{}
	n := 0
	for _, s := range ss {
		n += len(s)
		index = append(index, 0)
	}

	const maxInt = int(^uint(0) >> 1)
	for len(res) < n {
		minVal := maxInt
		minIndex := 0
		for i := range ss {
			if index[i] == len(ss[i]) {
				continue
			}
			if minVal > ss[i][index[i]] {
				minVal = ss[i][index[i]]
				minIndex = i
			}
		}
		index[minIndex]++
		res = append(res, minVal)
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func sortConcurrent(s []int) {
	k := 4 // number of goroutines
	n := len(s)
	m := n / k // number of integers that each goroutine need to sort
	if n%k > 0 {
		m++
	}
	m = max(m, 1) // each goroutine need to sort at least one integer

	ss := [][]int{}

	for i := 0; i < k; i++ {
		start := i * m
		end := min(i*m+m, n)

		if start >= n {
			ss = append(ss, []int{})
		} else {
			ss = append(ss, s[start:end])
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(k)
	for i := 0; i < k; i++ {
		go sortLinear(ss[i], &wg, i)
	}
	wg.Wait()

	fmt.Printf("\nAll goroutines finished, start merging sorted slices\n")

	temp := merge(ss)
	for i := range s {
		s[i] = temp[i]
	}

	fmt.Printf("\nMerging finished\n")
}

func main() {
	s := []int{}
	x := 0
	fmt.Println("Please enter some integers separated by space (Enter x to end the input):")
	fmt.Print("> ")
	for {
		nRead, err := fmt.Scan(&x)
		if nRead != 1 || err != nil {
			break
		}
		s = append(s, x)
	}
	fmt.Printf("\nOriginal slice:\n")
	fmt.Println(s)

	sortConcurrent(s)

	fmt.Printf("\nSorted slice:\n")
	fmt.Println(s)
}
