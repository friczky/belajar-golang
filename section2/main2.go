package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving")
}

func (c Car) GetName() string {
	return c.Name
}

func (c Car) GetAge() int {
	return c.Age
}

func (c Car) GetModelNo() int {
	return c.ModelNo
}

func main() {
	c := Car{
		Name:    "BMW",
		Age:     2,
		ModelNo: 2,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
	fmt.Println(c.GetAge())
	fmt.Println(c.GetModelNo())

}
