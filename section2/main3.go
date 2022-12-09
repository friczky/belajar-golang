package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

type BMW struct {
	BMWModel string
}

func (l *Lambo) Stop() {
	fmt.Println("Lambo is stopped")
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo is driving")
	fmt.Println(l.LamboModel)
}

func (b *BMW) Drive() {
	fmt.Println("BMW is driving")
	fmt.Println(b.BMWModel)
}

func main() {
	l := Lambo{"Lambo"}
	b := BMW{"BMW"}
	l.Drive()
	b.Drive()
}
