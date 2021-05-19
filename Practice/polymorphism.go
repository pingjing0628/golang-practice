package main

import (
	"fmt"
	"sort"
)

type AnimalIF interface {
	Sleep()
	Age() int
	Type() string
}
type Animal struct {
	MaxAge int
}

/*=======================================================*/

type Cat struct {
	Animal Animal
}

func (this *Cat) Sleep() {
	fmt.Println("Cat need sleep")
}
func (this *Cat) Age() int {
	return this.Animal.MaxAge
}
func (this *Cat) Type() string {
	return "Cat"
}

/*=======================================================*/

type Dog struct {
	Animal Animal
}

func (this *Dog) Sleep() {
	fmt.Println("Dog need sleep")
}
func (this *Dog) Age() int {
	return this.Animal.MaxAge
}
func (this *Dog) Type() string {
	return "Dog"
}

/*=======================================================*/

func Factory(name string) AnimalIF {
	switch name {
	case "dog":
		return &Dog{Animal{MaxAge: 20}}
	case "cat":
		return &Cat{Animal{MaxAge: 10}}
	default:
		panic("No such animal")
	}
}

/*======================================================*/

func main() {
	animal := Factory("dog")
	animal.Sleep()
	fmt.Printf("%s max age is: %d", animal.Type(), animal.Age())
}
