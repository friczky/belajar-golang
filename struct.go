package main

import "fmt"

type person struct {
	name, country string
	age           int
}

func main() {
	var a person
	a.name = "Fadilah Riczky"
	a.country = "Japan"
	a.age = 23

	fmt.Println(a)

	b := person{
		name:    "Fadilah Riczky",
		country: "Japan",
		age:     23,
	}

	fmt.Println(b)
}
