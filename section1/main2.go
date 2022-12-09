package main

import "fmt"

func todo() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	arr2 = append(arr2, "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z")
	fmt.Println(arr, arr2)
}

func main() {
	todo()
}
