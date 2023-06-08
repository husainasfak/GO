package main

import "fmt"

type Animal interface {
	says() string
	numberOfLegs() int
}
type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German shephered",
	}

	PrintInfo(&dog);

	gorilla := Gorilla{
		Name:"Jock",
		Color: "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.says(), "and has ", a.numberOfLegs(), "legs")
}

func (d *Dog) says() string{
	return "woof"
}

func (d *Dog)  numberOfLegs() int {
	return 4;
}

func (d *Gorilla) says() string{
	return "Ugh"
}

func (d *Gorilla)  numberOfLegs() int {
	return 4;
}