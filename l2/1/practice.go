package main

import "fmt"

type Human struct {
	name string
	age  int
	sex  string
}

type Action struct {
	Human
	effect string
}

func (h *Human) changeName(newName string) {
	h.name = newName
}

func main() {
	jora := Action{
		Human: Human{
			name: "jora",
			age:  20,
			sex:  "Q",
		},
		effect: "umer",
	}
	jora.changeName("jija")
	fmt.Printf("ya teper: %s, i ya: %s", jora.name, jora.effect)
}
