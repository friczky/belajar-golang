package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(5, 5))

	result, err := sqrtroot(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sum(x int, y int) int {
	z := x + y

	return z
}

func sqrtroot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undifined for negative number")
	}
	return math.Sqrt(x), nil
}
