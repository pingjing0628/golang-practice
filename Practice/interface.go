package main

import (
	"fmt"
)

type Peopleee interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() Peopleee {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
