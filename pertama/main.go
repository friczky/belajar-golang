package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hallo Belajar Golang")
	// sentence := TestAja()

	// fmt.Println(sentence)
	result := calculation.Add(10, 10)

	fmt.Println(result)
}
