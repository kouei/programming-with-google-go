package main

import "fmt"

func main() {
	var x float64
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Error!")
	}
	xTrunc := int64(x)
	fmt.Println(xTrunc)
}
