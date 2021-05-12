package main

import (
	"fmt"
)

type Peoplee interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo Peoplee = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
