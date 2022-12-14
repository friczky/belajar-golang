package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Hello World")
	Anything(2.3)
	Anything("Hello World")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap['name']= "John"
	mymap['age']= 23
	fmt.Println(mymap)
}
