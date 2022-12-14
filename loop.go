package main

import "fmt"

func main() {

	// print 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	name := []string{"alfin", "jumadi", "fadilah", "riczky"}

	// range loop slice
	for index, value := range name {
		fmt.Println("index:", index, "value:", value)
	}

	// range loop map
	stock := map[string]int{
		"apple":  10,
		"orange": 20,
		"grape":  30,
	}

	for key, value := range stock {
		fmt.Println("key:", key, "value:", value)
	}
}
