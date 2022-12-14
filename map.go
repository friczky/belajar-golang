package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["long"] = 1
	x["wide"] = 2
	x["high"] = 3

	delete(x, "long")

	fmt.Println(x)

	y := map[string]string{
		"firstName": "Fadilah",
		"lastName":  "Riczky",
	}

	fmt.Println(y)

}
