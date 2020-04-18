package main

import (
	"fmt"
	"os"
)

// Name stores first name and last name
type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Please enter the input file:")
	inputFile := ""
	fmt.Scan(&inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	names := []Name{}
	for {
		name := Name{}
		nRead, err := fmt.Fscanf(f, "%20s %20s", &name.fname, &name.lname)

		if nRead != 2 || err != nil {
			break
		}

		names = append(names, name)
	}

	fmt.Println("\nAll names:")
	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
