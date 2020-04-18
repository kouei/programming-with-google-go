package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func readStr() string {
	reader := bufio.NewReader(os.Stdin)
	var s string
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s = s[:len(s)-1]
	return s
}

func printMap(m map[string]string) {
	for k, v := range m {
		println(k, ":", v)
	}
}

func main() {
	fmt.Println("Please enter your name:")
	name := readStr()
	fmt.Println("Please enter your address:")
	address := readStr()

	m := map[string]string{
		"name":    name,
		"address": address,
	}

	fmt.Println("\nOriginal map:")
	printMap(m)

	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mm := make(map[string]string)
	json.Unmarshal(jsonBytes, &mm)

	fmt.Println("\nMap from unmarshalled json:")
	printMap(mm)

	fmt.Println("\nJson:")
	fmt.Println(string(jsonBytes))
}
