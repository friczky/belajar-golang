package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2)

	fmt.Println(x)

	y := []int{1, 2, 3, 4, 5}

	fmt.Println(y)
}
