package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isFound(s string) bool {
	s = strings.ToLower(s)
	return strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n")
}

func main() {
	fmt.Println("Please enter a string:")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s = s[:len(s)-1]

	if isFound(s) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
