package main

import "fmt"

func main() {
	var x [5]int
	x[2] = 10

	fmt.Println(x)

	y := [5]int{1, 2, 3, 4, 5}

	fmt.Println(y)
}
